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

func Transform(obj map[string]json.RawMessage) map[string]interface{} {

    var (
        _string string
        _int int
        _float64 float64
        err error
    )

    // TODO: if type byte, transform to map[string]json.RawMessage
    //if bytes, ok := obj.(byte); ok == true {
    //}

    _obj := make(map[string]json.RawMessage)

    result := make(map[string]interface{})

    for key, val := range obj {
        for {
            err = json.Unmarshal(val, &_string)
            if (err == nil) {
                result[key] = _string
                break
            }
            err = json.Unmarshal(val, &_int)
            if (err == nil) {
                result[key] = _int
                break
            }
            err = json.Unmarshal(val, &_float64)
            if (err == nil) {
                result[key] = _float64
                break
            }
            err = json.Unmarshal(val, &_obj)
            if (err == nil) {
                tmp := Transform(_obj)
                result[key] = tmp
                break
            }
            log.Fatal("Syntax error")
        }
    }

    return result
}
