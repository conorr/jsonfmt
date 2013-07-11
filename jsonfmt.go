package main

import (
    "bytes"
    "encoding/json"
    "io"
    "io/ioutil"
    "os"
    "log"
    "regexp"
    "fmt"
)

func main() {

    if (len(os.Args) < 2) {
        fmt.Println("Usage: jsonfmt <json-file>");
        return
    }

    filename := os.Args[1]

    in, err := ioutil.ReadFile(filename)
    if (err != nil) {
        log.Fatal(err)
        return
    }

    out := bytes.NewBufferString("")

    if err := json.Indent(out, in, "", "\t"); err == nil {
        io.Copy(os.Stdout, out)
    } else {
        log.Fatal(err)
    }
}

func isJSONP(b []byte) bool {
    re := "([a-zA-Z_0-9\\.]*\\()|(\\);?$)"
    isMatch, _ := regexp.Match(re, b);
    return isMatch
}
