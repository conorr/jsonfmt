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
    contents, err := ioutil.ReadFile(filename)
    if (err != nil) {
        log.Fatal(err)
        return
    }

    // check if it's jsonp and print accordlingly
    isJSONP, parts := parseJSONP(contents)
    if (!isJSONP) {
        // regular json
        printFormattedJSON(contents)
    } else {
        // jsonp
        fmt.Print(string(parts[1]))
        printFormattedJSON(parts[2])
        fmt.Print(string(parts[3]) + "\n")
    }
}

func printFormattedJSON(jsonstr []byte) {
    out := bytes.NewBufferString("")
    if err := json.Indent(out, jsonstr, "", "\t"); err == nil {
        io.Copy(os.Stdout, out)
    } else {
        log.Fatal(err)
    }
}

func parseJSONP(str []byte) (bool, [][]byte) {
    re, _ := regexp.Compile(`^([A-Za-z_0-9\.]+\()(.*)(\))$`)
    matches := re.FindAllSubmatch(str, -1)
    if (matches == nil) {
        return false, nil
    } else if (len(matches[0]) != 4) {
        log.Fatal("Bad JSON!")
    }
    return true, matches[0]
}
