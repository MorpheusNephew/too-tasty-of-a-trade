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

	watchlistData, err := ttClient.GetPublicWatchlist("tasty IVR", "Equity")

	if err != nil {
		fmt.Println("Something bad happened getting watchlist data", err)
	}

	firstTenEntries := []string{}

	firstTenSlice := watchlistData.Data.WatchlistEntries[:10]

	for _, entry := range firstTenSlice {
		firstTenEntries = append(firstTenEntries, entry.Symbol)
	}

	// Get market metrics
	marketMetricsResponseBody, err := ttClient.GetMarketMetrics(firstTenEntries)

	if err != nil {
		fmt.Println("So what had happened", err)
	}

	fmt.Println("All of the magic", marketMetricsResponseBody)

	firstItem := watchlistData.Data.WatchlistEntries[0]

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
