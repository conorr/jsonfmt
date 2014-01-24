package main

import (
    "bytes"
    "fmt"
    "sort"
)

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
    var delim string

    if lvl == 0 {
        indentp = ""
    } else {
        indentp = indentMultiple(lvl - 1)
    }
    
    fmt.Printf("%s{\n", indentp)

    keys := GetKeysArray(src, true)

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
            Indent(dst, _map, lvl + 1, "    ")
        }
    }
    fmt.Printf("%s}\n", indentp)

    return nil
}

func GetKeysArray(obj map[string]interface{}, sortKeys bool) []string {
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
