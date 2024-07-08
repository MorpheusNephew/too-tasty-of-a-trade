package tastyworks

import (
	"fmt"
	"strings"
)

type EarningsInfo struct {
	ExpectedReportDate string `json:"expected-report-date"`
}

type MarketMetricInfo struct {
	Symbol                string       `json:"symbol"`
	ImpliedVolatilityRank string       `json:"implied-volatility-index-rank"`
	Earnings              EarningsInfo `json:"earnings"`
}

type MarketMetricInfoItems struct {
	Items []MarketMetricInfo `json:"items"`
}

type MarketMetricInfoResponse struct {
	Data MarketMetricInfoItems `json:"data"`
}

var marketMetricsUrl = fmt.Sprintf("%s/market-metrics", baseUrl)

func (t *TTClient) GetMarketMetrics(symbols []string) (MarketMetricInfoResponse, error) {
	url := fmt.Sprintf("%s?symbols=%s", marketMetricsUrl, strings.Join(symbols, ","))

	returnError := func(err error) (MarketMetricInfoResponse, error) {
		return MarketMetricInfoResponse{}, err
	}

	resp, err := t.get(url)

	if err != nil {
		return returnError(err)
	}

	responseBody := MarketMetricInfoResponse{}

	err = convertResponseToJson(*resp, &responseBody)

	if err != nil {
		return returnError(err)
	}

	return responseBody, nil
}
