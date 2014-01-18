package main

import (
    "os"
    "io"
    "fmt"
    "log"
    "bytes"
    "regexp"
    "errors"
    "encoding/json"
)

func main() {

    var (
        head bytes.Buffer
        body bytes.Buffer
        tail bytes.Buffer
    )

    // Parse args.
    if (len(os.Args) < 2) {
        fmt.Println("Usage: jsonfmt <json-file>");
        return
    }
    filename := os.Args[1]

    // Open file and read into buffer.
    fi, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, 1024)
    for {
        n, err := fi.Read(data)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
        body.Write(data[:n])
    }
    fi.Close()

    // Try parsing JSONP.
    if parts, err := parseJSONP(body.Bytes()); err == nil {
        head.Write(parts[0])
        body.Reset()
        body.Write(parts[1])
        tail.Write(parts[2])
    }

    // Make a new buffer of indented JSON.
    bodyIndented := bytes.NewBufferString("")
    err = json.Indent(bodyIndented, body.Bytes(), "", "    ")
    if err != nil {
        if serr, ok := err.(*json.SyntaxError); ok {
            fmt.Printf("Syntax error at byte %d: %s\n", serr.Offset, serr.Error())
        } else {
            fmt.Println(err)
        }
        os.Exit(1)
    }

    // Write the buffer into the same file.
    fo, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    fo.Write(head.Bytes())
    fo.Write(bodyIndented.Bytes())
    fo.Write(tail.Bytes())
    fo.Close()
}

func parseJSONP(contents []byte) ([][]byte, error) {
    re, _ := regexp.Compile("^([A-Za-z_0-9.]+[(]{1})(.*)([)]|[)][\n]+)$")
    matches := re.FindAllSubmatch(contents, -1)
    if len(matches) == 0 {
        fmt.Println("case 1")
        return nil, errors.New("Could not parse into JSONP")
    }
    parts := matches[0]
    if len(parts) < 3 {
        fmt.Println("case 2")
        return nil, errors.New("Could not parse into JSONP")
    }
    return parts[1:], nil
}
