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

func Writef(dst *bytes.Buffer, format string, a ...interface{}) {
    str := fmt.Sprintf(format, a...)
    dst.WriteString(str)
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
        Writef(dst, "\"%s\"", _str)
    } else if _int, ok := src.(int); ok {
        Writef(dst, "%d", _int)
    } else if _float64, ok := src.(float64); ok {
        Writef(dst, "%v", _float64)
    } else if _bool, ok := src.(bool); ok {
        Writef(dst, "%v", _bool)
    } else if _arr, ok := src.([]interface{}); ok {

        Writef(dst, "[\n")
        final := len(_arr) - 1
        for i, item := range _arr {
            Writef(dst, "%s", makeIndent(depth + 1))
            indent(dst, item, indentStr, depth + 1)
            if i != final {
                Writef(dst, ",")
            }
            Writef(dst, "\n")
        }

        Writef(dst, "%s]", makeIndent(depth))

    } else if _map, ok := src.(map[string]interface{}); ok {

        Writef(dst, "{\n")

        keys := getKeysArray(_map, false)
        final := len(keys) - 1
        for i, key := range keys {
            Writef(dst, "%s\"%s\": ", makeIndent(depth + 1), key)
            indent(dst, _map[key], indentStr, depth + 1)
            if i != final {
                Writef(dst, ",")
            }
            Writef(dst, "\n")
        }

        Writef(dst, "%s}", makeIndent(depth))

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
