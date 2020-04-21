package flags

import (
	"fmt"
	"reflect"
	"strings"
)

func SetFlagByStruct(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	farrt := t.Elem()
	farrv := v.Elem()
	fieldnum := farrt.NumField()
	for i := 0; i < fieldnum; i++ {
		fieldt := farrt.Field(i)
		fieldv := farrv.Field(i)
		lowername := strings.ToLower(fieldt.Name)
		uppername := strings.ToUpper(fieldt.Name)
		fmt.Println(fmt.Sprint(fieldv.Interface()))
		switch fieldt.Type.Kind() {
		case reflect.Slice:
			tmp := fieldv.Interface().([]string)
			fmt.Println(strings.Join(tmp, ","))
			SetFlag(lowername, uppername, "", strings.Join(tmp, ","))
		default:
			SetFlag(lowername, uppername, "", fmt.Sprint(fieldv.Interface()))
		}
	}
}

func LoadStruct(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	farrt := t.Elem()
	farrv := v.Elem()
	fieldnum := farrt.NumField()
	for i := 0; i < fieldnum; i++ {
		fieldt := farrt.Field(i)
		fieldv := farrv.Field(i)
		lowername := strings.ToLower(fieldt.Name)
		// uppername := strings.ToUpper(fieldt.Name)
		switch fieldt.Type.Kind() {
		case reflect.Float64:
			fieldv.SetFloat(GetFlagByFloat64(lowername))
			// 	flag.Float64Var(&(fieldv.Interface().(float64)), lowername, 0, lowername)
		case reflect.Int:
			fieldv.SetInt(int64(GetFlagByInt(lowername)))
			// flag.IntVar(&(fieldv.Interface().(int)), lowername, 0, lowername)
		case reflect.String:
			fieldv.SetString(GetFlagByString(lowername))
		case reflect.Slice:
			fieldv.Set(reflect.ValueOf(GetFlagBySlice(lowername)))
		case reflect.Bool:
			fieldv.Set(reflect.ValueOf(GetFlagByBool(lowername)))
		default:
			GetFlagByString(lowername)
			// 	flag.StringVar(&(fieldv.Interface().(string)), lowername, "", lowername)
			// case reflect.Bool:
			// 	flag.BoolVar(&(fieldv.Interface().(bool)), lowername, false, lowername)
		}
	}
}
