package tastyworks

import "fmt"

var marketMetricsUrl = fmt.Sprintf("%s/market-metrics", baseUrl)

func Something() {
	fmt.Println("Market Metrics", marketMetricsUrl)
}
