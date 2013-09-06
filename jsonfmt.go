package main

import (
    "os"
    "io"
    "fmt"
    "log"
    "bytes"
    "regexp"
    "encoding/json"
)

func main() {

    var (
        buf bytes.Buffer
        jsonBody []byte
        isJSONP bool
        jsonpParts [][]byte
    )

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

    // check if it's JSONP
    if isJSONP, jsonpParts := checkJSONP(buf.Bytes()); isJSONP {
        jsonBody = jsonpParts[1]
    } else {
        jsonBody = buf.Bytes()
    }

    // make a new buffer of indented json
    cleanJSON := bytes.NewBufferString("")
    err = json.Indent(cleanJSON, jsonBody, "", "    ")
    if err != nil {
        if serr, ok := err.(*json.SyntaxError); ok {
            fmt.Printf("Syntax error at byte %d: %s\n", serr.Offset, serr.Error())
        } else {
            fmt.Println(err)
        }
        os.Exit(1)
    }

    if isJSONP {
        // add the head and tail padding
        jsonpParts[1] = cleanJSON.Bytes()
        cleanJSON = bytes.NewBuffer(bytes.Join(jsonpParts, []byte{}))
    }

    fmt.Println(cleanJSON.String())

    /*
    // write the buffer into the same file
    fo, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    fo.Write(cleanJSON.Bytes())
    fo.Close()
    */
}

func checkJSONP(contents []byte) (bool, [][]byte) {
    re, _ := regexp.Compile(`^([A-Za-z_0-9\.]+\()(.*)(\))$`)
    matches := re.FindAllSubmatch(contents, -1)
    fmt.Println(matches)
    os.Exit(1)
    if (matches == nil) {
        return false, nil
    } else if (len(matches[0]) != 4) {
        log.Fatal("Bad JSON!")
    }
    return true, matches[0]
}
