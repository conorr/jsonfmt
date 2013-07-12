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

    // parse args
    if (len(os.Args) < 2) {
        fmt.Println("Usage: jsonfmt <json-file>");
        return
    }
    filename := os.Args[1]

    // read file
    in, err := ioutil.ReadFile(filename)
    if (err != nil) {
        log.Fatal(err)
        return
    }

    isJSONP := parseJSONP(in)

    if (isJSONP == nil) {
        printJSON(in)
    } else {
        fmt.Print(string(isJSONP[1]))
        printJSON(isJSONP[2])
        fmt.Print(string(isJSONP[3]))
    }
    fmt.Print("\n")

}

func printJSON(jsonstr []byte) {
    out := bytes.NewBufferString("")
    if err := json.Indent(out, jsonstr, "", "\t"); err == nil {
        io.Copy(os.Stdout, out)
    } else {
        log.Fatal(err)
    }
}

func parseJSONP(str []byte) [][]byte {
    re, _ := regexp.Compile(`^([A-Za-z_0-9\.]+\()(.*)(\))$`)
    matches := re.FindAllSubmatch(str, -1)
    if (matches == nil) {
        return nil
    } else if (len(matches[0]) != 4) {
        log.Fatal("Bad JSON!")
    }
    return matches[0]
}
