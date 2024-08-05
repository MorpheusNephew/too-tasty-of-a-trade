package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	expectedNumOfArgs := 2

	if len(argsWithoutProg) != expectedNumOfArgs {
		fmt.Printf("Program should have %d arguments\n", expectedNumOfArgs)
		os.Exit(1)
	}

	strikePrices := []float64{}

	for _, val := range argsWithoutProg {
		if convertedFloat, err := strconv.ParseFloat(val, 64); err == nil {
			strikePrices = append(strikePrices, convertedFloat)
		} else {
			fmt.Printf("Can only submit float values: %s", val)
			os.Exit(1)
		}
	}

	strikeDiff := math.Abs(strikePrices[0] - strikePrices[1])

	minimumAmountPerDollarWide := 0.33
	maximumAmountPerDollarWide := 0.44

	minimumProfitPerContract := minimumAmountPerDollarWide * strikeDiff
	maximumProfitPerContract := maximumAmountPerDollarWide * strikeDiff

	fmt.Printf("The minimum profit is %.2f per contract or $%.2f total\nThe maximum profit is %.2f per contract or $%.2f total\n",
		minimumProfitPerContract,
		minimumProfitPerContract*100,
		maximumProfitPerContract,
		maximumProfitPerContract*100,
	)
}
