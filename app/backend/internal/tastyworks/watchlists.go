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

/*
Get public watchlist with a specific watchlist name will return a list of symbols on that list.
The information here that will be used for the first iteration of this project is `implied-volatility-index-rank`.
Anything greater than 0.60 should be added to a list for further inspection
*/
func (t *TTClient) GetPublicWatchList(listName, instrumentType string) (*WatchListResponse, error) {
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

	matchedInstrumentEntries := []WatchListEntry{}

	for _, watchlistEntry := range responseBody.Data.WatchListEntries {
		if watchlistEntry.InstrumentType == instrumentType {
			matchedInstrumentEntries = append(matchedInstrumentEntries, watchlistEntry)
		}
	}

	responseBody.Data.WatchListEntries = matchedInstrumentEntries

	return &responseBody, nil
}
