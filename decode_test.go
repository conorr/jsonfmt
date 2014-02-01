package main

import (
    "testing"
    //"fmt"
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

func TestRawInterfaceMapArray(t *testing.T) {

    assertEqualInt := func(x int, y int) {
        if x != y {
            t.Errorf("Expected %d, got %d", x, y)
        }
    }

    test := "{\"foo\":[1,2,3]}"
    result, _ := RawInterfaceMap([]byte(test))

    arr, ok := result["foo"].([]interface{})
    if ok == false {
        t.Errorf("oooh no!")
    }

    tmp := make([]int, len(arr))
    for i := 0; i < len(arr); i++ {
        tmp[i] = arr[i].(int)
    }

    assertEqualInt(tmp[0], 1)
    assertEqualInt(tmp[1], 2)
    assertEqualInt(tmp[2], 3)

}

func TestRawInterfaceMapArray2(t *testing.T) {

    assertEqualInt := func(x int, y int) {
        if x != y {
            t.Errorf("Expected %d, got %d", x, y)
        }
    }

    test := "{\"foo\":{\"bar\":[1,2,3]}}"
    result, _ := RawInterfaceMap([]byte(test))

    arr1, ok := result["foo"].(map[string]interface{})
    if ok == false {
        t.Errorf("oooh no!")
    }

    arr2, ok := arr1["bar"].([]interface{})
    if ok == false {
        t.Errorf("oooh no!")
    }

    tmp := make([]int, len(arr2))
    for i := 0; i < len(arr2); i++ {
        tmp[i] = arr2[i].(int)
    }

    assertEqualInt(tmp[0], 1)
    assertEqualInt(tmp[1], 2)
    assertEqualInt(tmp[2], 3)

}

func TestRawInterfaceMapErrors(t *testing.T) {

    tests := []string{
        "{\"foo\": \"bar}",
        "{\"foo\": [1, 2 3]}",
    }

    for _, test := range tests {
        _, err := RawInterfaceMap([]byte(test))
        if err == nil {
            t.Errorf("Expected error due to bad syntax")
        }
    }
}
