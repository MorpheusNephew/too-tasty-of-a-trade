package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MorpheusNephew/ttoat/v2/internal/tastyworks"
)

func main() {
	fmt.Println("Welcome to Too Tasty of a Trade!")

	above90WatchlistName := "above90-autogenerated"

	above80WatchlistName := "above80-autogenerated"

	above70WatchlistName := "above70-autogenerated"

	httpClient := &http.Client{}

	ttClient := tastyworks.GetTTClient(httpClient)

	username, password := "", ""

	ttClient.CreateSession(username, password)

	watchlistData, err := ttClient.GetPublicWatchlist("tasty IVR", "Equity")

	if err != nil {
		fmt.Println("Something bad happened getting watchlist data", err)
	}

	allSymbols := []string{}

	for _, entry := range watchlistData.Data.WatchlistEntries {
		allSymbols = append(allSymbols, entry.Symbol)
	}

	fmt.Printf("The length of all of the symbols is %d\n", len(allSymbols))

	// Get market metrics
	marketMetricsResponseBody, err := ttClient.GetMarketMetrics(allSymbols)

	if err != nil {
		fmt.Println("So what had happened", err)
	}

	marketMetricsItems := marketMetricsResponseBody.Data.Items

	fmt.Printf("The length of market data is %d\n", len(marketMetricsItems))

	above90, above80, above70 := []tastyworks.WatchlistEntry{}, []tastyworks.WatchlistEntry{}, []tastyworks.WatchlistEntry{}

	sixtyDaysFromNow := time.Now().Add(time.Hour * 24 * 60)

	fmt.Printf("60 days from now? %d-%d-%d\n", sixtyDaysFromNow.Year(), sixtyDaysFromNow.Month(), sixtyDaysFromNow.Day())

	for _, item := range marketMetricsItems {
		impliedVolatilityRank, err := strconv.ParseFloat(item.ImpliedVolatilityRank, 64)

		if err != nil {
			continue
		}

		dateSlice := strings.Split(item.Earnings.ExpectedReportDate, "-")

		earningsYear, _ := strconv.ParseInt(dateSlice[0], 10, 64)
		earningsMonth, _ := strconv.ParseInt(dateSlice[1], 10, 64)
		earningsDay, _ := strconv.ParseInt(dateSlice[2], 10, 64)

		earningsDate := time.Date(int(earningsYear), time.Month(earningsMonth), int(earningsDay), 0, 0, 0, 0, time.Local)

		if earningsDate.Compare(sixtyDaysFromNow) == -1 {
			fmt.Println("Earnings", earningsDate, "is too close")
			continue
		}

		symbol := item.Symbol

		if impliedVolatilityRank >= 0.9 {
			above90 = append(above90, tastyworks.WatchlistEntry{Symbol: symbol, InstrumentType: "Equity"})
			fmt.Println("Above 90 earnings date", item.Earnings.ExpectedReportDate)
		} else if impliedVolatilityRank >= 0.8 {
			above80 = append(above80, tastyworks.WatchlistEntry{Symbol: symbol, InstrumentType: "Equity"})
			fmt.Println("Above 80 earnings date", item.Earnings.ExpectedReportDate)
		} else if impliedVolatilityRank >= 0.7 {
			above70 = append(above70, tastyworks.WatchlistEntry{Symbol: symbol, InstrumentType: "Equity"})
			fmt.Println("Above 70 earnings date", item.Earnings.ExpectedReportDate)
		}
	}

	fmt.Println(len(above90), len(above80), len(above70))

	// firstTwo := watchlistData.Data.WatchlistEntries[:2]

	ttClient.UpdatePrivateWatchlist(above90WatchlistName, &above90)
	ttClient.UpdatePrivateWatchlist(above80WatchlistName, &above80)
	ttClient.UpdatePrivateWatchlist(above70WatchlistName, &above70)

	// ttClient.DeletePrivateWatchlist(watchlistName)

	ttClient.RemoveSession()

	// r := gin.Default()

	// r.GET("/health-check", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "You've made it",
	// 	})
	// })

	// r.Run(":3000")
}
