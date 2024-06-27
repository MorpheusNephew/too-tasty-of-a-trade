package tastyworks

import (
	"fmt"
	"strings"
)

type MarketMetricInfo struct {
	Symbol                string `json:"symbol"`
	ImpliedVolatilityRank string `json:"implied-volatility-index-rank"`
}

type MarketMetricInfoItems struct {
	Items []MarketMetricInfo `json:"items"`
}

type MarketMetricInfoResponse struct {
	Data MarketMetricInfoItems `json:"data"`
}

var marketMetricsUrl = fmt.Sprintf("%s/market-metrics", baseUrl)

func (t *TTClient) GetMarketMetrics(symbols []string) (*MarketMetricInfoResponse, error) {
	url := fmt.Sprintf("%s?symbols=%s", marketMetricsUrl, strings.Join(symbols, ","))

	resp, err := t.get(url)

	if err != nil {
		return nil, err
	}

	responseBody := &MarketMetricInfoResponse{}

	err = convertResponseToJson(resp, responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
