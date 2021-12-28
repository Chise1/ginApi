package models

import (
	"reflect"
)

func StructToMap(n interface{}) map[string]interface{} {
	t := reflect.TypeOf(n)
	v := reflect.ValueOf(n)
	switch n.(type) {
	case IOpenapiStruct:
		return n.(IOpenapiStruct).Map()
	default:

	}
	res := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsZero() {
			//fmt.Println(t.Field(i).Name, "--")
			continue
		}
		fieldName := t.Field(i).Tag.Get("json")
		if t.Field(i).Type.Kind() == reflect.Struct {
			switch v.Field(i).Interface().(type) {
			case IOpenapiStruct:
				res[fieldName] = v.Field(i).Interface().(IOpenapiStruct).Map()
			default:
				res[fieldName] = StructToMap(v.Field(i).Interface())
			}
		} else if t.Field(i).Type.Kind() == reflect.Map {
			field := v.Field(i)
			iter := field.MapRange()
			mapValue := map[string]interface{}{}
			for iter.Next() {
				k := iter.Key().String()
				value := iter.Value()
				if value.Kind() == reflect.Struct {
					mapValue[k] = StructToMap(value.Interface())
				} else {
					mapValue[k] = value.Interface()
				}
			}
			res[fieldName] = mapValue
		} else if t.Field(i).Type.Kind() == reflect.Slice {
			var list []interface{}
			slice := v.Field(i)
			l := slice.Len()
			for j := 0; j < l; j++ {
				a := slice.Index(j)
				if a.Kind() == reflect.Struct {
					list = append(list, StructToMap(a.Interface()))
				} else {
					list = append(list, slice.Index(j).Interface())
				}
			}
			res[fieldName] = list
		} else if t.Field(i).Type.Kind() == reflect.Ptr {
			res[fieldName] = StructToMap(reflect.Indirect(v.Field(i)).Interface())

		} else {
			res[fieldName] = v.Field(i).Interface()
		}
	}
	return res
}
