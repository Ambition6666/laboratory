package utils

import "reflect"

// 将结构体所需字段导出成数组, 使用标签"excel"
func StoArr(s any) []any {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	
	// 如果是指针类型将会转换成对应值类型
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}

	// 如果不是结构体类型, 将其会返回空数组
	if t.Kind() != reflect.Struct {
		return nil
	}

	res := make([]any, 0)
	for i := 0; i < t.NumField(); i++{
		f := t.Field(i)
		if (f.Tag.Get("excel") != "") {
			if f.Type.Kind() == reflect.Struct || f.Type.Kind() == reflect.Pointer {
				res = append(res, StoArr(v.Field(i).Interface())...)
			} else {
				res = append(res, v.Field(i).Interface())
			}
		}
	} 

	return res
}