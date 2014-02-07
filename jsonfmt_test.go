package main

import (
	"bytes"
	"errors"
	"fmt"
	"jsonfmt/util"
	"testing"
)

func TestJSONFmt(t *testing.T) {

	testBuf := util.ReadFile("testfiles/1_test.json")
	resultBuf := JSONFmt(testBuf, false)
	expectBuf := util.ReadFile("testfiles/1_expect.json")

	err := compareBuffers(expectBuf, resultBuf)
	if err != nil {
		t.Error(err)
	}
}

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

func compareBuffers(expectBuf *bytes.Buffer, resultBuf *bytes.Buffer) error {

	// Trim newlines. TODO: determine if newlines should matter.
	expectBytes := bytes.TrimRight(expectBuf.Bytes(), "\n")
	resultBytes := bytes.TrimRight(resultBuf.Bytes(), "\n")

	for i := 0; i < len(expectBytes); i++ {
		if i > (len(resultBytes) - 1) {
			if err := dumpBuffer("testfiles/_result.json", resultBuf); err != nil {
				fmt.Println(err)
			}
			return errors.New("Result doesn't match expected byte length!")
		}
		if expectBytes[i] != resultBytes[i] {
			if err := dumpBuffer("testfiles/.result.json", resultBuf); err != nil {
				fmt.Println(err)
			}
			str := fmt.Sprintf("Expected \"%s\", got \"%s\" at byte %d",
				string(expectBytes[i]), string(resultBytes[i]), i)
			return errors.New(str)
		}
	}

	return nil
}

func dumpBuffer(filename string, buf *bytes.Buffer) error {
	err := util.WriteFile(filename, buf);
	if err != nil {
		return err
	}
	return nil
}
