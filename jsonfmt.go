package main

import (
    "os"
    "io"
    "fmt"
    "log"
    "bytes"
    "encoding/json"
)

func main() {

    // parse args
    if (len(os.Args) < 2) {
        fmt.Println("Usage: jsonfmt <json-file>");
        return
    }
    filename := os.Args[1]

    // open file and read it into buffer
    fi, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, 1024)
    var buf bytes.Buffer
    for {
        n, err := fi.Read(data)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
        buf.Write(data[:n])
    }
    fi.Close()

    // make a new buffer of indented json
    cleanJSON := bytes.NewBufferString("")
    err = json.Indent(cleanJSON, buf.Bytes(), "", "    ")
    if err != nil {
        if serr, ok := err.(*json.SyntaxError); ok {
            fmt.Printf("Syntax error at byte %d: %s\n", serr.Offset, serr.Error())
        } else {
            fmt.Println(err)
        }
        os.Exit(1)
    }

    // write the buffer into the same file
    fo, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    fo.Write(cleanJSON.Bytes())
    fo.Close()
}
