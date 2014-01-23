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

func RawInterfaceMap(bytes []byte) (map[string]interface{}, error) {
    obj := make(map[string]json.RawMessage)
    err := json.Unmarshal(bytes, &obj)
    if err != nil {
        return nil, err
    }
    result, err := DecodeRawMessageMap(obj)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DecodeRawMessageMap(obj map[string]json.RawMessage) (map[string]interface{}, error) {

    var (
        _string string
        _int int
        _float64 float64
        err error
    )

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
                tmp, err := DecodeRawMessageMap(_obj)
                if err != nil {
                    return nil, err
                }
                result[key] = tmp
                break
            }
            // TODO: return different error; SyntaxError has an offset
            return nil, &json.SyntaxError{}
        }
    }

    return result, nil
}

func Indent(dst *bytes.Buffer, src map[string]interface{}, lvl int, indent string) error {

    // TODO: eliminate lvl in favor of a closure

    indentMultiple := func(l int) string {
        var result string
        for i := 0; i < (l + 1); i++ {
            result += indent
        }
        return result
    }

    indentn := indentMultiple(lvl)
    var indentp string
    if lvl == 0 {
        indentp = ""
    } else {
        indentp = indentMultiple(lvl - 1)
    }
    
    fmt.Printf("%s{\n", indentp)
    for key, val := range src {

        if _string, ok := val.(string); ok {
            fmt.Printf("%s\"%s\": \"%s\"\n", indentn, key, _string)
        } else if _int, ok := val.(int); ok {
            fmt.Printf("%s\"%s\": %d\n", indentn, key, _int)
        } else if _float64, ok := val.(float64); ok {
            fmt.Printf("%s\"%s\": %v\n", indentn, key, _float64)
        } else if _map, ok := val.(map[string]interface{}); ok {
            Indent(dst, _map, lvl + 1, "    ")
        }
    }
    fmt.Printf("%s}\n", indentp)

    return nil
}
