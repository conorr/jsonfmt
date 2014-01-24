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

func indent(dst *bytes.Buffer, src map[string]interface{}, indentStr string, lvl int) error {

    indentLvl := func(lvl int) string {
        if lvl == -1 {
            return ""
        }
        str := ""
        for i := 0; i < (lvl + 1); i++ {
         str += indentStr
        }
        return str
    }

    del := "" 
    fmt.Printf("{")
    keys := getKeysArray(src, true)
    for _, key := range keys {

        fmt.Printf("%s\n%s", del, indentLvl(lvl))

        val := src[key]

        if _string, ok := val.(string); ok {
            fmt.Printf("\"%s\": \"%s\"", key, _string)
        } else if _int, ok := val.(int); ok {
            fmt.Printf("\"%s\": %d", key, _int)
        } else if _float64, ok := val.(float64); ok {
            fmt.Printf("\"%s\": %v", key, _float64)
        } else if _map, ok := val.(map[string]interface{}); ok {
            indent(dst, _map, "    ", lvl + 1)
        } else {
            log.Fatal()
        }

        del = ","
    }
    fmt.Printf("\n%s}", indentLvl(lvl - 1))

    if lvl == 0 {
        fmt.Printf("\n")
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
