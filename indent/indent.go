package indent

import (
	"bytes"
	"fmt"
	"log"
	"sort"
)

type BufferWriter struct {
	buf *bytes.Buffer
}

func (writer BufferWriter) Writef(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	writer.buf.WriteString(str)
}

func Indent(dst *bytes.Buffer, src interface{}, indentStr string, sortKeys bool) error {
	return indent(dst, src, indentStr, 0, sortKeys)
}

func indent(dst *bytes.Buffer, src interface{}, indentStr string, depth int, sortKeys bool) error {

	writer := BufferWriter{buf: dst}

	makeIndent := func(depth int) string {
		str := ""
		for i := 0; i < depth; i++ {
			str += indentStr
		}
		return str
	}

	if _str, ok := src.(string); ok {
		writer.Writef("%q", _str)
	} else if _int, ok := src.(int); ok {
		writer.Writef("%d", _int)
	} else if _float64, ok := src.(float64); ok {
		writer.Writef("%v", _float64)
	} else if _bool, ok := src.(bool); ok {
		writer.Writef("%v", _bool)
	} else if _arr, ok := src.([]interface{}); ok {

		writer.Writef("[\n")
		final := len(_arr) - 1
		for i, item := range _arr {
			writer.Writef("%s", makeIndent(depth+1))
			indent(dst, item, indentStr, depth+1, sortKeys)
			if i != final {
				writer.Writef(",")
			}
			writer.Writef("\n")
		}

		writer.Writef("%s]", makeIndent(depth))

	} else if _map, ok := src.(map[string]interface{}); ok {

		writer.Writef("{\n")

		keys := getKeysArray(_map, sortKeys)
		final := len(keys) - 1
		for i, key := range keys {
			writer.Writef("%s%q: ", makeIndent(depth+1), key)
			indent(dst, _map[key], indentStr, depth+1, sortKeys)
			if i != final {
				writer.Writef(",")
			}
			writer.Writef("\n")
		}

		writer.Writef("%s}", makeIndent(depth))

	} else {
		return errors.New("Could not process interface")
	}

	return nil
}

// Given a map of type map[string]interface{}, return an array of its keys.
// If sortKeys is true, sort the keys alphabetically.
func getKeysArray(obj map[string]interface{}, sortKeys bool) []string {
	arr := make([]string, len(obj))
	i := 0
	for key, _ := range obj {
		arr[i] = key
		i++
	}
	if sortKeys {
		sort.Strings(arr)
	}
	return arr
}
