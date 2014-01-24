package main

import (
    "testing"
    "bytes"
    //"fmt"
)



func TestParseJSONP1(t *testing.T) {

    var buf bytes.Buffer
    var err error

    buf.Write([]byte("CALLBACK({\"foo\":\"bar\"})"))
    _, err = ParseJSONP(buf.Bytes())
    if err != nil {
        t.Error("Valid JSONP throwing error")
    }
}

func TestParseJSONP2(t *testing.T) {

    var buf bytes.Buffer
    var err error

    buf.Write([]byte("CALLBACK{\"foo\":\"bar\"})"))
    _, err = ParseJSONP(buf.Bytes())
    if err == nil {
        t.Error("Invalid JSONP not throwing error")
    }
}

func TestParseJSONP3(t *testing.T) {

    tests := make(map[string][]string)

    // Tests
    tests["CALLBACK({})"] = []string{"CALLBACK(", "{}", ")"}
    tests["\nCALLBACK({})\n"] = []string{"\nCALLBACK(", "{}", ")\n"}
    tests["\nCall.back({})\n"] = []string{"\nCall.back(", "{}", ")\n"}

    for test, expect := range tests {
        result, _ := ParseJSONP([]byte(test))
        for i := 0; i < 3; i++ {
            if string(result[i]) != expect[i] {
                t.Errorf("Expected %s, got %s", expect[i], string(result[i]))
            }
        }
    }
}

func TestRawInterfaceMap(t *testing.T) {

    // Tests
    tests := make(map[string]interface{})
    tests["{\"foo\":\"bar\"}"] = "bar"
    tests["{\"foo\":7}"] = 7
    tests["{\"foo\":3.14}"] = 3.14

    for test, expect := range tests {
        result, _ := RawInterfaceMap([]byte(test))
        if result["foo"] != expect {
            t.Errorf("Expected %s, got %s", expect, result["foo"])
        }
    }

}

func TestRawInterfaceMapErrors(t *testing.T) {

    test := []byte("{\"foo\":\"bar}")
    _, err := RawInterfaceMap(test)

    if err == nil {
        t.Errorf("Expected error due to bad syntax")
    }

}

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
