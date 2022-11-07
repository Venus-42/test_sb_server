package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("static"))
	//r.StaticFile("/")
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello, %s", name)
	})
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := context.Request.Method + " => " + name + " is " + action
		context.String(http.StatusOK, message)
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
