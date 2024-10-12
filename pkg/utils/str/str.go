package str

import (
	"reflect"
)

func ForEach(str any, fn func(key string, val any)) {
	value := reflect.ValueOf(str)

	for i := range value.NumField() {
		key := value.Type().Field(i).Name
		val := value.Field(i)

		if !val.IsValid() || IsEmptyValue(val) {
			continue
		}

		if val.Kind() == reflect.Ptr {
			if !val.IsNil() {
				fn(key, val.Interface())
			}
		} else {
			fn(key, val.Interface())
		}
	}
}

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Map:
		return v.IsNil() || v.Len() == 0
	case reflect.Ptr:
		return v.IsNil()
	}
	return false
}
