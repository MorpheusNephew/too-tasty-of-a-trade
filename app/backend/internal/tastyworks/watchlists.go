package tastyworks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type WatchListEntry struct {
	Symbol         string `json:"symbol"`
	InstrumentType string `json:"instrument_type"`
}

type WatchList struct {
	Name             string           `json:"name"`
	WatchListEntries []WatchListEntry `json:"watchlist-entries"`
}

type WatchLists struct {
	Items []WatchList `json:"items"`
}

type WatchListsResponse struct {
	Context string     `json:"context"`
	Data    WatchLists `json:"data"`
}

type WatchListResponse struct {
	Context string    `json:"context"`
	Data    WatchList `json:"data"`
}

func (t *TTClient) GetPublicWatchLists() {
	url := fmt.Sprintf("%s/public-watchlists", baseUrl)

	resp, err := t.get(url)

	if err != nil {
		panic("Getting public watchlists is busted")
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	responseBody := WatchListsResponse{}

	err = json.Unmarshal(bodyBytes, &responseBody)

	if err != nil {
		panic("An error occurred in the unmarshalling")
	}

	fmt.Println("What are the results?", responseBody)
}

func (t *TTClient) GetPublicWatchList(listName string) {
	url := fmt.Sprintf("%s/public-watchlists/%s", baseUrl, url.PathEscape((listName)))

	fmt.Println("What is this url?", url)

	resp, err := t.get(url)

	if err != nil {
		panic("Getting public watchlists is busted")
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	responseBody := WatchListResponse{}

	err = json.Unmarshal(bodyBytes, &responseBody)

	if err != nil {
		panic("An error occurred in the unmarshalling")
	}

	fmt.Println("What are the results?", responseBody)
}
