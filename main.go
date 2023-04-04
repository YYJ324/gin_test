package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	//curl local:8080/hello
	//返回json
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "index",
		})
	})
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "world",
		})
	})
	err := engine.Run(":1234")
	if err != nil {
		log.Fatalln(err)
	}

}
