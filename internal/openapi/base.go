package openapi

type IOpenapiStruct interface {
	Map() map[string]interface{}
}
type IReference interface {
	GetRef() struct{}
}
