package tastyworks

import "net/http"

const baseUrl string = "https://api.tastyworks.com"

func GetTTClient() TTClient {
	httpClient := http.Client{}

	// Create TTClient which will be used to make all TastyTrade requests
	return TTClient{HttpClient: httpClient}
}
