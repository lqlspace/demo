package reflectutil

import (
	"fmt"
	"reflect"
)


func GetFnNameByReflect(fn interface{}) (string, error) {
	t := reflect.TypeOf(fn)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("invalid parameter: expected struct, acutal %s", t.Kind().String())
	}

	return t.Name(), nil 
}