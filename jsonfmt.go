package main

import (
    "bytes"
    "encoding/json"
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

    // check if it's jsonp and handle accordingly
    if check, parts := isJSONP(contents); check == true {
        // handle jsonp
        fmt.Print(string(parts[1]))
        printFormattedJSON(parts[2])
        fmt.Print(string(parts[3]) + "\n")
    } else {
        // handle regular json
        printFormattedJSON(contents)
    }
}

func printFormattedJSON(contents []byte) {
    formatted := bytes.NewBufferString("")
    if err := json.Indent(formatted, contents, "", "\t"); err == nil {
        fmt.Print(formatted.String())
    } else {
        log.Fatal(err)
    }
}

func isJSONP(contents []byte) (bool, [][]byte) {
    re, _ := regexp.Compile(`^([A-Za-z_0-9\.]+\()(.*)(\))$`)
    matches := re.FindAllSubmatch(contents, -1)
    if (matches == nil) {
        return false, nil
    } else if (len(matches[0]) != 4) {
        log.Fatal("Bad JSON!")
    }
    return true, matches[0]
}
