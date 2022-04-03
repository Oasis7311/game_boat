package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func MarshalIgnoreErr(v interface{}) string {
	if v == nil {
		return ""
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func, reflect.UnsafePointer:
		if reflect.ValueOf(v).IsNil() {
			return ""
		}
	}

	b, err := json.Marshal(v)
	if err != nil {
		return "ERR_JSON_MARSHAL"
	}
	return string(b)
}

func JsonStrFormatIgnoreErr(jsonStr interface{}) string {
	str := MarshalIgnoreErr(jsonStr)

	formatStr := new(bytes.Buffer)
	_ = json.Indent(formatStr, []byte(str), "", "\t")

	return formatStr.String()
}
