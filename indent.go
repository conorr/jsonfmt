package main

import (
    "bytes"
    "fmt"
    "sort"
    "log"
)

func Indent(dst *bytes.Buffer, src map[string]interface{}, indentStr string) error {
    return indent(dst, src, indentStr, 0)
}

func indent(dst *bytes.Buffer, src interface{}, indentStr string, depth int) error {
    makeIndent := func(depth int) string {
        str := ""
        for i := 0; i < depth; i++ {
            str += indentStr
        }
        return str
    }

    if _str, ok := src.(string); ok {
        fmt.Printf("\"%s\"", _str)
    } else if _int, ok := src.(int); ok {
        fmt.Printf("%d", _int)
    } else if _float64, ok := src.(float64); ok {
        fmt.Printf("%v", _float64)
    } else if _bool, ok := src.(bool); ok {
        fmt.Printf("%v", _bool)
    } else if _arr, ok := src.([]interface{}); ok {

        fmt.Printf("[\n")
        final := len(_arr) - 1
        for i, item := range _arr {
            fmt.Printf("%s", makeIndent(depth + 1))
            indent(dst, item, indentStr, depth + 1)
            if i != final {
                fmt.Printf(",")
            }
            fmt.Printf("\n")
        }

        fmt.Printf("%s]", makeIndent(depth))

    } else if _map, ok := src.(map[string]interface{}); ok {

        fmt.Printf("{\n")

        keys := getKeysArray(_map, false)
        final := len(keys) - 1
        for i, key := range keys {
            fmt.Printf("%s\"%s\": ", makeIndent(depth + 1), key)
            indent(dst, _map[key], indentStr, depth + 1)
            if i != final {
                fmt.Printf(",")
            }
            fmt.Printf("\n")
        }

        fmt.Printf("%s}", makeIndent(depth))

    } else {
        log.Fatal("Don't know what to do with it!")
    }

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
