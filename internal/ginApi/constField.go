package ginApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type HttpMethod string

const Get HttpMethod = "GET"
const Post HttpMethod = "POST"

type Null struct {
}
type ErrInfo struct {
	status int
	data   string
	c      *gin.Context
}

func (n *ErrInfo) Ret() {
	n.c.String(n.status, n.data)
}

type StringRes struct {
	status int
	data   string
	c      *gin.Context
}

func (n StringRes) Dec() {
	n.c.String(n.status, n.data)
}
func NewStringRes(c *gin.Context, status int, data string) StringRes {
	if status == 0 {
		status = 200
	}
	return StringRes{status: status, data: data, c: c}
}

type JsonRes struct {
	status int
	data   interface{}
	c      *gin.Context
}

func (n JsonRes) Dec() {
	n.c.JSON(n.status, n.data)
}
func NewJsonRes(c *gin.Context, status int, data interface{}) JsonRes {
	if status == 0 {
		status = 200
	}
	return JsonRes{status: status, data: data, c: c}
}

type Field struct {
	notRepeated bool
	name        string
	FieldType   reflect.Type
	place       string
}

func Iter(f interface{}) func(ctx *gin.Context) {
	a := reflect.TypeOf(f)
	inD := a.In(1)
	var intDict = []Field{}
	inl := inD.NumField()
	for i := 0; i < inl; i++ {
		inF := inD.Field(i)
		tags := inF.Tag.Get("gA")
		if tags == "" {
			intDict = append(intDict, Field{
				name:      inF.Name,
				FieldType: inF.Type,
			})
		} else if tags == "param" {
			intDict = append(intDict, Field{
				name:      inF.Name,
				FieldType: inF.Type,
				place:     "param",
			})
		} else {
			//tagL:=strings.Split(tags,";")
			//for _,tag:=range tagL{
			//
			//}
			intDict = append(intDict, Field{
				name:      inF.Name,
				FieldType: inF.Type,
			})
		}
	}
	return func(ctx *gin.Context) {
		resErros := StringRes{}
		res := reflect.New(inD)
		for i, fieldInfo := range intDict {
			if fieldInfo.place == "param" {
				r, ok := ctx.Params.Get(fieldInfo.name)
				if !ok {
					resErros.status = 400
				} else {
					res.Elem().Field(i).Set(reflect.ValueOf(r))
				}
			} else {
				r, ok := ctx.Get(fieldInfo.name)
				if !ok {
					resErros.status = 400
				} else {
					res.Elem().Field(i).Set(reflect.ValueOf(r))
				}
			}
		}
		if resErros.status != 0 {
			ctx.JSON(400, resErros.data)
			return
		}
		var inFields []reflect.Value
		inFields = append(inFields, reflect.ValueOf(ctx), reflect.Indirect(res))
		ret := reflect.ValueOf(f).Call(inFields)
		IretData, IerrInfo := ret[0].Interface(), ret[1].Interface()
		errInfo := IerrInfo.(*ErrInfo)
		if errInfo != nil {
			errInfo.Ret()
		} else {
			switch IretData.(type) {
			case JsonRes:
				retData := IretData.(JsonRes)
				retData.Dec()
			default:
				ctx.JSON(200, IretData)
			}
		}
	}
}
func Docs(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "docs.tmpl", gin.H{})
	})
	r.GET("/openapi.json", func(context *gin.Context) {
		context.String(200, "{\"openapi\":\"3.0.2\",\"info\":{\"title\":\"FastAPI\",\"version\":\"0.1.0\"},\"paths\":{\"/user/{Name}\":{\"get\":{\"summary\":\"Read Root\",\"operationId\":\"read_root_user__Name__get\",\"parameters\":[{\"required\":true,\"schema\":{\"title\":\"Name\",\"type\":\"string\"},\"name\":\"Name\",\"in\":\"path\"}],\"responses\":{\"200\":{\"description\":\"Successful Response\",\"content\":{\"application/json\":{\"schema\":{}}}},\"422\":{\"description\":\"Validation Error\",\"content\":{\"application/json\":{\"schema\":{\"$ref\":\"#/components/schemas/HTTPValidationError\"}}}}}}}},\"components\":{\"schemas\":{\"HTTPValidationError\":{\"title\":\"HTTPValidationError\",\"type\":\"object\",\"properties\":{\"detail\":{\"title\":\"Detail\",\"type\":\"array\",\"items\":{\"$ref\":\"#/components/schemas/ValidationError\"}}}},\"ValidationError\":{\"title\":\"ValidationError\",\"required\":[\"loc\",\"msg\",\"type\"],\"type\":\"object\",\"properties\":{\"loc\":{\"title\":\"Location\",\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"msg\":{\"title\":\"Message\",\"type\":\"string\"},\"type\":{\"title\":\"Error Type\",\"type\":\"string\"}}}}}}")
	})
}
