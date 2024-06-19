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

type Watchlist struct {
	Name             string           `json:"name"`
	WatchListEntries []WatchListEntry `json:"watchlist-entries"`
}

type Watchlists struct {
	Items []Watchlist `json:"items"`
}

type WatchListsResponse struct {
	Context string     `json:"context"`
	Data    Watchlists `json:"data"`
}

type WatchListResponse struct {
	Context string    `json:"context"`
	Data    Watchlist `json:"data"`
}

var watchlistsUrl = fmt.Sprintf("%s/public-watchlists", baseUrl)

func (t *TTClient) GetPublicWatchLists() (*WatchListsResponse, error) {
	resp, err := t.get(watchlistsUrl)

	if err != nil {
		return nil, err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	responseBody := WatchListsResponse{}

	err = json.Unmarshal(bodyBytes, &responseBody)

	if err != nil {
		return nil, err
	}

	fmt.Println("What are the results?", responseBody)

	return &responseBody, nil
}

func (t *TTClient) GetPublicWatchList(listName string) (*WatchListResponse, error) {
	url := fmt.Sprintf("%s/%s", watchlistsUrl, url.PathEscape((listName)))

	resp, err := t.get(url)

	if err != nil {
		return nil, err
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	responseBody := WatchListResponse{}

	err = json.Unmarshal(bodyBytes, &responseBody)

	if err != nil {
		return nil, err
	}

	fmt.Println("What are the results?", responseBody)
	return &responseBody, nil
}
