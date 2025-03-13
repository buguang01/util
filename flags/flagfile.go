package flags

import (
	"fmt"
	"reflect"
	"strings"
)

// 设置结构体为启动参数
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
		flag := fieldt.Tag.Get("flag")
		// flagarr := strings.Split(flag, ",")
		// node := ""
		// defval := ""
		// for _, v := range flagarr {
		// 	if strings.HasPrefix(v, "default=") {
		// 		defval = strings.TrimPrefix(v, "default=")
		// 		continue
		// 	} else {
		// 		node = v
		// 	}
		// }
		// fmt.Println(fmt.Sprint(fieldv.Interface()))
		switch fieldt.Type.Kind() {
		case reflect.Slice:
			tmp := fieldv.Interface().([]string)
			// fmt.Println(strings.Join(tmp, ","))
			SetFlag(lowername, uppername, flag, strings.Join(tmp, ","))
		default:
			SetFlag(lowername, uppername, flag, fmt.Sprint(fieldv.Interface()))
		}
	}
}

// 启动参数写进结构体中
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
		switch fieldt.Type.Kind() {
		case reflect.Float64:
			fieldv.SetFloat(GetFlagByFloat64(lowername))
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int8:
			fieldv.SetInt(int64(GetFlagByInt(lowername)))
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint8:
			fieldv.SetUint(uint64(GetFlagByUint(lowername)))
		case reflect.String:
			fieldv.SetString(GetFlagByString(lowername))
		case reflect.Slice:
			fieldv.Set(reflect.ValueOf(GetFlagBySlice(lowername)))
		case reflect.Bool:
			fieldv.Set(reflect.ValueOf(GetFlagByBool(lowername)))
		default:
			GetFlagByString(lowername)

		}
	}
}
