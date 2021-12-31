package utils

import "github.com/ginApi/openapi/models"

// Struct2Example 把reqStruct和repStruct转为example方便openapi显示
func Struct2Example(n interface{}) models.Example {
	return models.Example{}
}
