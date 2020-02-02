package util

import (
	"errors"
	"reflect"
)

//反射创建新对象。
func ReflectNew(target interface{}) (interface{}, error) {
	if target == nil {
		return nil, errors.New("反射创建新对象时，参数不能未空。")
	}
	t := reflect.TypeOf(target)
	if t.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		t = t.Elem()
	}

	newStruc := reflect.New(t) // 调用反射创建对象
	return newStruc.Interface(), nil
}
