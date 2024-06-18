package client

import (
	"io"
	"net/http"
)

type TTClient struct {
	HttpClient   *http.Client
	SessionToken string
}

var baseUrl = "https://api.tastyworks.com"

var contentType = "application/json"

var userAgent = "tastytrade-api-client/1.0"

func GetTTClient() *TTClient {
	// Create http client to be used for underlying requests
	httpClient := &http.Client{}

	// Create TTClient which will be used to make all TastyTrade requests
	return &TTClient{HttpClient: httpClient}
}

func (t *TTClient) addHeaders(req *http.Request) {
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", userAgent)
}

func (t *TTClient) addAuthHeaders(req *http.Request) {
	t.addHeaders(req)
	req.Header.Set("Authorization", t.SessionToken)
}

func (t *TTClient) get(url string) (resp *http.Response, err error) {
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	t.addAuthHeaders(request)

	return t.HttpClient.Do(request)
}

func (t *TTClient) post(url string, body io.Reader) (resp *http.Response, err error) {
	request, _ := http.NewRequest(http.MethodPost, url, body)

	t.addHeaders(request)

	return t.HttpClient.Do(request)
}

func (t *TTClient) delete(url string) (resp *http.Response, err error) {
	request, _ := http.NewRequest(http.MethodDelete, url, nil)

	t.addAuthHeaders(request)

	return t.HttpClient.Do(request)
}
