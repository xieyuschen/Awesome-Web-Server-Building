package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "listen port")
	flag.Parse()
	g := gin.Default()
	g.GET("m1", M1, M2, func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"message": "m1",
		})
	})

	g.GET("m1NonNext", M1NonNext, M2NonNext, func(context *gin.Context) {
		context.JSON(200, map[string]string{
			"message": "m1NonNext",
		})
	})

	_ = g.Run(":" + *port)
}

func M1(ctx *gin.Context) {
	fmt.Println("M1 in")
	ctx.Next()
	fmt.Println("M1 out")
}
func M2(ctx *gin.Context) {
	fmt.Println("M2 in")
	ctx.Next()
	fmt.Println("M2 out")
}
func M1NonNext(ctx *gin.Context) {
	fmt.Println("M1NonNext in")
	fmt.Println("M1NonNext out")
}
func M2NonNext(ctx *gin.Context) {
	fmt.Println("M2NonNext in")
	ctx.Next()
	fmt.Println("M2NonNext out")
}
