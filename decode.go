package main

import (
    "encoding/json"
)

func RawInterfaceMap(bytes []byte) (map[string]interface{}, error) {
    obj := make(map[string]json.RawMessage)
    err := json.Unmarshal(bytes, &obj)
    if err != nil {
        return nil, err
    }
    result, err := DecodeRawMessageMap(obj)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func DecodeRawMessageMap(obj map[string]json.RawMessage) (map[string]interface{}, error) {

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
                tmp, err := DecodeRawMessageMap(_obj)
                if err != nil {
                    return nil, err
                }
                result[key] = tmp
                break
            }
            // TODO: return different error; SyntaxError has an offset
            return nil, &json.SyntaxError{}
        }
    }

    return result, nil
}
