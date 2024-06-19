package main

import (
	"fmt"

	"github.com/MorpheusNephew/ttoat/v2/internal/tastyworks"
)

func main() {
	fmt.Println("Welcome to Too Tasty of a Trade!")

	// r := gin.Default()

	// r.GET("/health-check", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "You've made it",
	// 	})
	// })

	// r.Run(":3000")

	ttClient := tastyworks.GetTTClient()

	username, password := "", ""

	ttClient.CreateSession(username, password)

	ttClient.GetPublicWatchLists()

	ttClient.GetPublicWatchList("tasty IVR")

	ttClient.RemoveSession()
}
