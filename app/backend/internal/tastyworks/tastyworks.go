package tastyworks

import "net/http"

const baseUrl string = "https://api.tastyworks.com"

func GetTTClient(httpClient *http.Client) *TTClient {
	// Create TTClient which will be used to make all TastyTrade requests
	return &TTClient{HttpClient: httpClient}
}
