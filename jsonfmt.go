package main

import (
    "os"
    "io"
    "fmt"
    "log"
    "bytes"
    "regexp"
    "errors"
    "jsonfmt/decode"
    "jsonfmt/indent"
)

const READBYTES int = 1024
const JSONP_RE string = "^([\n]?[A-Za-z_0-9.]+[(]{1})(.*)([)]|[)][\n]+)$"

func main() {

    var (
        head bytes.Buffer
        body bytes.Buffer
        tail bytes.Buffer
    )

    // Parse args.
    if (len(os.Args) < 2) {
        fmt.Println("Usage: jsonfmt [file]");
        os.Exit(1)
    }
    filename := os.Args[1]

    // Open file and read into buffer.
    fi, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, READBYTES)
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
    if parts, err := ParseJSONP(body.Bytes()); err == nil {
        head.Write(parts[0])
        body.Reset()
        body.Write(parts[1])
        tail.Write(parts[2])
    }

    // Make a new buffer of indented JSON.
    indentedBody := bytes.NewBufferString("")
    i, err := decode.RawInterfaceMap(body.Bytes())
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    indent.Indent(indentedBody, i, "    ", false)

    // Write the buffer into the same file.
    fo, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    fo.Write(head.Bytes())
    fo.Write(indentedBody.Bytes())
    fo.Write(tail.Bytes())
    fo.Close()
}

func ParseJSONP(contents []byte) ([][]byte, error) {
    re, _ := regexp.Compile(JSONP_RE)
    matches := re.FindAllSubmatch(contents, -1)
    if len(matches) == 0 {
        return nil, errors.New("Could not parse into JSONP")
    }
    parts := matches[0]
    if len(parts) < 3 {
        return nil, errors.New("Could not parse into JSONP")
    }
    return parts[1:], nil
}
