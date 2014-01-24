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

    b := make(map[string]interface{})
    b["foo"] = "bar"
    i["maptest"] = b

    Indent(&buf, i, 0, "    ")
    
}

func TestGetKeysArray(t *testing.T) {

    i := make(map[string]interface{})
    i["pineapple"] = "bar"
    i["banana"] = 2
    i["apple"] = 3.14

    arr := GetKeysArray(i, true)

    if arr[0] != "apple" {
        t.Errorf("oh no!")
    }

    if arr[1] != "banana" {
        t.Errorf("oh no!")
    }

    if arr[2] != "pineapple" {
        t.Errorf("oh no!")
    }
}
