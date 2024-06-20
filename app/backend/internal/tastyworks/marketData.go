package tastyworks

import (
	"fmt"
	"strings"
)

type OptionExpirationImpliedVolatility struct {
	ExpirationDate    string `json:"expiration-date"`
	SettlementType    string `json:"settlement-type"`
	OptionChainType   string `json:"option-chain-type"`
	ImpliedVolatility string `json:"implied-volatility"`
}

type MarketMetricInfo struct {
	Symbol                              string                              `json:"symbol"`
	ImpliedVolatilityIndex              float64                             `json:"implied-volatility-index"`
	ImpliedVolatilityIndex5DayChange    float64                             `json:"implied-volatility-index-5-day-change"`
	ImpliedVolatilityRank               float64                             `json:"implied-volatility-rank"`
	ImpliedVolatilityPercentile         float64                             `json:"implied-volatility-percentile"`
	Liquidity                           float64                             `json:"liquidity"`
	LiquidityRank                       float64                             `json:"liquidity-rank"`
	LiquidityRating                     int32                               `json:"liquidity-rating"`
	OptionExpirationImpliedVolatilities []OptionExpirationImpliedVolatility `json:"option-expiration-implied-volatilities"`
}

type MarketMetricInfoItems struct {
	Items []MarketMetricInfo `json:"items"`
}

type MarketMetricInfoData struct {
	Data MarketMetricInfoItems `json:"data"`
}

type MarketMetricInfoResponse struct {
	Data MarketMetricInfoData `json:"data"`
}

var marketMetricsUrl = fmt.Sprintf("%s/market-metrics", baseUrl)

func (t *TTClient) GetMarketMetrics(symbols []string) (*MarketMetricInfoResponse, error) {
	url := fmt.Sprintf("%s?symbols=%s", marketMetricsUrl, strings.Join(symbols, ","))

	resp, err := t.get(url)

	if err != nil {
		return nil, err
	}

	responseBody := MarketMetricInfoResponse{}

	err = convertResponseToJson(resp, &responseBody)

	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}
