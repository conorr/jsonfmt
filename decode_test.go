package main

import (
    "testing"
)

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
