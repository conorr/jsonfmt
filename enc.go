package main

import "encoding/json"
import "fmt"
import "log"

func main() {
    jsonStr := []byte("{\"foo\":2,\"bar\":\"barbar\",\"floaty\":5.0}")

    obj := make(map[string]json.RawMessage)
    err := json.Unmarshal(jsonStr, &obj)
    if err != nil {
        log.Fatal(err)
    }

    result := make(map[string]interface{})

    for key, val := range obj {

        var (
            _string string
            _int int
            _float64 float64
            err error
        )

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

            log.Fatal("couldn't do anything with value!")

        }

    }

    fmt.Println(result)

}
