package main

import (
    "fmt"
    "log"
    "encoding/json"
)

func main() {

    jsonStr := []byte("{\"foo\":2,\"bar\":\"barbar\",\"floaty\":5.0,\"x\":{\"a\":5,\"b\":{\"foo\":\"bar\"}}}")

    obj := make(map[string]json.RawMessage)
    err := json.Unmarshal(jsonStr, &obj)
    if err != nil {
        log.Fatal(err)
    }

    result := Transform(obj)
    fmt.Println(result)
}

func Transform(obj map[string]json.RawMessage) map[string]interface{} {

    var (
        _string string
        _int int
        _float64 float64
        err error
    )

    _obj := make(map[string]json.RawMessage)

    result := make(map[string]interface{})

    for key, val := range obj {
        for {

            err = json.Unmarshal(val, &_string)
            if (err == nil) {
                result[key] = _string
                break
            }

            err = json.Unmarshal(val, &_int)
            if (err == nil) {
                result[key] = _int
                break
            }

            err = json.Unmarshal(val, &_float64)
            if (err == nil) {
                result[key] = _float64
                break
            }

            err = json.Unmarshal(val, &_obj)
            if (err == nil) {
                tmp := Transform(_obj)
                result[key] = tmp
                break
            }

            log.Fatal("Couldn't do anything with value!")
        }
    }

    return result

}
