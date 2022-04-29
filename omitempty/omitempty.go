package omitempty

import (
	"encoding/json"
	"reflect"
	"strings"
)

func MarshalJSON(obj interface{}) ([]byte, error) {
	fieldHasOmitEmpty := func(field reflect.StructField) bool {
		// Get the field tag value
		tag := field.Tag.Get("json")
		options := strings.Split(tag, ",")
		for _, option := range options {
			if option == "omitempty" {
				return true
			}
		}
		return false
	}

	st := reflect.TypeOf(obj)
	fs := []reflect.StructField{}
	for i := 0; i < st.NumField(); i++ {
		fs = append(fs, st.Field(i))
	}

	for i, _ := range fs {
		if !fieldHasOmitEmpty(fs[i]) {
			continue
		}

		val := reflect.ValueOf(obj)
		valueField := val.Field(i)
		f := valueField.Interface()

		if f == nil {
			continue
		}

		switch reflect.TypeOf(f).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			continue
		}

		if isempty, ok := f.(interface {
			IsEmpty() bool
		}); ok {
			if !isempty.IsEmpty() {
				continue
			}
			fs[i].Tag = reflect.StructTag(`json:"-"`)
		}
		continue
	}

	st2 := reflect.StructOf(fs)

	v := reflect.ValueOf(obj)
	v2 := v.Convert(st2)
	return json.Marshal(v2.Interface())
}
