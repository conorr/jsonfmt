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

    var head bytes.Buffer
    var body bytes.Buffer

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
        fmt.Println("is jsonp!")
        head.Write(parts[0])
        body.Reset()
        body.Write(parts[1])
    }

    // Make a new buffer of indented JSON.
    cleanJSON := bytes.NewBufferString("")

    err = json.Indent(cleanJSON, body.Bytes(), "", "    ")
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
    // If JSONP, add padding back in.
    // TODO: need the if?
    if len(head.Bytes()) > 0 {
        fo.Write(head.Bytes())
        //fo.Write([]byte("("))
    }
    fo.Write(cleanJSON.Bytes())
    //if len(head.Bytes()) > 0 {
        //fo.Write([]byte(")\n"))
    //}
    fo.Close()
}

func parseJSONP(contents []byte) ([][]byte, error) {
    //s := string(contents)
    //fmt.Println(s)
    re, _ := regexp.Compile("^([A-Za-z_0-9.]+)[(](.*)[)]([\n]|)$")
    matches := re.FindAllSubmatch(contents, -1)
    if len(matches) == 0 {
        return nil, errors.New("Could not parse into JSONP")
    }
    parts := matches[0]
    if len(parts) < 3 {
        return nil, errors.New("Could not parse into JSONP")
    }
    //return nil, nil
    return parts[1:], nil
}
