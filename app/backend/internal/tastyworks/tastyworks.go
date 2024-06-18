package tastyworks

import "net/http"

func GetTTClient() *TTClient {
	// Create http client to be used for underlying requests
	httpClient := &http.Client{}

	// Create TTClient which will be used to make all TastyTrade requests
	return &TTClient{HttpClient: httpClient}
}
