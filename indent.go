package main

import (
    "bytes"
    "fmt"
    "sort"
)

func Indent(dst *bytes.Buffer, src map[string]interface{}, indentStr string) error {
    return indent(dst, src, indentStr, 0)
}

func indent(dst *bytes.Buffer, src map[string]interface{}, indentStr string, lvl int) error {

    indentMultiple := func(l int) string {
        var result string
        for i := 0; i < (l + 1); i++ {
            result += indentStr
        }
        return result
    }

    indentn := indentMultiple(lvl)
    var indentp string
    var delim string

    if lvl == 0 {
        indentp = ""
    } else {
        indentp = indentMultiple(lvl - 1)
    }
    
    fmt.Printf("%s{\n", indentp)

    keys := getKeysArray(src, true)

    for i, key := range keys {

        val := src[key]

        if i == (len(keys) - 1) {
            delim = ""
        } else {
            delim = ","
        }

        if _string, ok := val.(string); ok {
            fmt.Printf("%s\"%s\": \"%s%s\"\n", indentn, key, _string, delim)
        } else if _int, ok := val.(int); ok {
            fmt.Printf("%s\"%s\": %d%s\n", indentn, key, _int, delim)
        } else if _float64, ok := val.(float64); ok {
            fmt.Printf("%s\"%s\": %v%s\n", indentn, key, _float64, delim)
        } else if _map, ok := val.(map[string]interface{}); ok {
            indent(dst, _map, "    ", lvl + 1)
        }
    }
    fmt.Printf("%s}\n", indentp)

    return nil
}

func getKeysArray(obj map[string]interface{}, sortKeys bool) []string {
    arr := make([]string, len(obj))
    i := 0
    for key, _ := range obj {
        arr[i] = key
        i++
    }
    if (sortKeys) {
        sort.Strings(arr)
    }
    return arr
}
