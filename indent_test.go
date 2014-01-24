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

    Indent(&buf, i, "    ")
    
}
