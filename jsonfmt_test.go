package main

import "testing"
import "bytes"

func TestParseJSONP(t *testing.T) {

    var buf bytes.Buffer
    var err error

    buf.Write([]byte("CALLBACK({\"foo\":\"bar\"})"))
    _, err = ParseJSONP(buf.Bytes())
    if err != nil {
        t.Fatal("Valid JSONP throwing error")
    }

    buf.Reset()

    buf.Write([]byte("CALLBACK{\"foo\":\"bar\"})"))
    _, err = ParseJSONP(buf.Bytes())
    if err == nil {
        t.Fatal("Invalid JSONP not throwing error")
    }

}
