package main

import (
	"fmt"
	"net/http"

	"github.com/MorpheusNephew/ttoat/v2/internal/tastyworks"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Too Tasty of a Trade!")

	ttClient := tastyworks.GetTTClient()

	username, password := "", ""

	ttClient.CreateSession(username, password)

	watchlistData, err := ttClient.GetPublicWatchList("tasty IVR", "Equity")

	if err != nil {
		fmt.Println("Something bad happened getting watchlist data", err)
	}

	firstItem := watchlistData.Data.WatchListEntries[0]

	ttClient.GetOptionChain(firstItem.Symbol)

	ttClient.RemoveSession()

	r := gin.Default()

	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "You've made it",
		})
	})

	r.Run(":3000")
}
