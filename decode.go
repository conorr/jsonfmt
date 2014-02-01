package main

import (
    "encoding/json"
    "errors"
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
    result := make(map[string]interface{})
    var err error
    for key, val := range obj {
        result[key], err = DecodeRawMessage(val)
        if err != nil {
            return nil, err
        }
    }
    return result, nil
}

func DecodeRawMessage(obj json.RawMessage) (interface{}, error) {

    var (
        _string string
        _int int
        _float64 float64
        _bool bool
        err error
    )

    _obj := make(map[string]json.RawMessage)
    _arr := []json.RawMessage{}
    var result interface{}

    for {
        err = json.Unmarshal(obj, &_string)
        if (err == nil) {
            result = _string
            break
        }

        err = json.Unmarshal(obj, &_int)
        if (err == nil) {
            result = _int
            break
        }

        err = json.Unmarshal(obj, &_float64)
        if (err == nil) {
            result = _float64
            break
        }

        err = json.Unmarshal(obj, &_bool)
        if (err == nil) {
            result = _bool
            break
        }

        err = json.Unmarshal(obj, &_arr)
        if err == nil {
            tmp := make([]interface{}, len(_arr))
            for i, el := range _arr {
                tmp[i], err = DecodeRawMessage(el)
                if err != nil {
                    return nil, err
                }
            }
            result = tmp
            break
        }

        err = json.Unmarshal(obj, &_obj)
        if (err == nil) {
            tmp, err := DecodeRawMessageMap(_obj)
            if err != nil {
                return nil, err
            }
            result = tmp
            break
        }

        return nil, errors.New("Syntax Error")
    }

    return result, nil
}
