package main

import (
    "bytes"
    "testing"
)

func TestIndent(t *testing.T) {

    var buf bytes.Buffer
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

    Indent(&buf, i, "    ")
}
