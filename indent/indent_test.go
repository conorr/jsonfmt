package indent

import (
    "bytes"
    "testing"
    "os"
    "log"
    "io"
    "fmt"
    "jsonfmt/decode"
)

func TestIndent(t *testing.T) {

    //var buf bytes.Buffer
    i := make(map[string]interface{})
    i["foo"] = "bar"
    i["num"] = 2
    i["floaty"] = 3.14
    i["bool"] = false

    b := make(map[string]interface{})
    b["foo"] = "bar"
    b["ima"] = "map"
    b["ican"] = 24
    i["maptest"] = b

    bc := make(map[string]interface{})
    bc["apples"] = "oranges"
    bc["bananas"] = "pineapples"
    i["fruits"] = bc

    arr := make([]interface{}, 3)
    arr[0] = bc
    arr[1] = bc
    arr[2] = bc
    i["arr"] = arr

    //Indent(&buf, i, "    ")
}

func TestIndentEndtoEnd(t *testing.T) {

    var bufIn bytes.Buffer
    var bufOut bytes.Buffer

    // Open file and read into buffer.
    fi, err := os.Open("../testfiles/test1.json")
    if err != nil {
        log.Fatal(err)
    }
    readBytes := make([]byte, 1024)
    for {
        n, err := fi.Read(readBytes)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if n == 0 {
            break
        }
        bufIn.Write(readBytes[:n])
    }
    fi.Close()

    obj, err := decode.RawInterfaceMap(bufIn.Bytes())
    if err != nil {
        t.Errorf("RawInterfaceMap returned error; possible syntax error")
        return
    }
    Indent(&bufOut, obj, "    ")

    // Trim newlines off right of file
    for index, expect := range bytes.TrimRight(bufIn.Bytes(), "\n") {
        if index > len(bufOut.Bytes()) + 1 {
            t.Errorf("bufOut was smaller than bufIn!")
            break
        }
        result := bufOut.Bytes()[index]
        if expect != result {
            t.Errorf("Expecting '%s', got '%s' at byte %d", string(expect),
                string(result), index)
            fmt.Println(string(bufOut.Bytes()))
            break
        }
    }
}

func TestWritef(t *testing.T) {

    var buf bytes.Buffer

    Writef(&buf, "hello")
    str := string(buf.Bytes())
    if str != "hello" {
        t.Errorf("Error!")
    }

    buf.Reset()
    
    Writef(&buf, "\"%s\"", "hello")
    str = string(buf.Bytes())
    if str != "\"hello\"" {
        t.Errorf("Error!")
    }

}
