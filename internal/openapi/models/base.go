package models

type IOpenapiStruct interface {
	Map() map[string]interface{}
}
type IReference interface {
	GetRef() struct{}
}
