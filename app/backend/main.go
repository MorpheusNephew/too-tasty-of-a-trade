package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Too Tasty of a Trade!")

	r := gin.Default()

	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You've made it",
		})
	})

	r.Run(":3000")
}
