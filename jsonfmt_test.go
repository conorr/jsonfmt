package main

import "testing"
import "bytes"
import "encoding/json"
import "log"
//import "fmt"

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

func TestTransform1(t *testing.T) {

    //jsonBytes := []byte("{\"foo\":2,\"bar\":\"barbar\",\"floaty\":5.0,\"x\":{\"a\":5,\"b\":{\"foo\":\"bar\"}}}")
    jsonBytes := []byte("{\"foo\":\"bar\"}")

    obj := make(map[string]json.RawMessage)
    err := json.Unmarshal(jsonBytes, &obj)
    if err != nil {
        log.Fatal(err)
    }

    result := Transform(obj)

    if result["foo"] != "bar" {
        t.Error()
    }
}

func TestTransform2(t *testing.T) {

    // Tests
    tests := make(map[string]interface{})
    tests["{\"foo\":\"bar\"}"] = "bar"
    tests["{\"foo\":7}"] = 7
    tests["{\"foo\":3.14}"] = 3.14

    for test, expect := range tests {
        obj := make(map[string]json.RawMessage)
        err := json.Unmarshal([]byte(test), &obj)
        if err != nil {
            t.Error()
        }
        result := Transform(obj)
        if result["foo"] != expect {
            t.Errorf("Expected %s, got %s", expect, result["foo"])
        }
    }

}
