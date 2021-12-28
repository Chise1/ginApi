package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ginApi/internal/ginApi"
)

type InUserName struct {
	Name string `gA:"param"`
}
type OutUserName struct {
	User  string
	Value string
}

func GetName(ctx *gin.Context, name InUserName) (OutUserName, *ginApi.ErrInfo) {
	return OutUserName{
		User: name.Name}, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user/:Name", ginApi.Iter(GetName))
	ginApi.Docs(r)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
