package str

import (
	"reflect"
)

func ForEach(str any, fn func(key string, val any)) {
	value := reflect.ValueOf(str)

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}

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
	case reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}
