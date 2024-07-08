package tastyworks

import (
	"fmt"
	"net/url"
)

type WatchlistEntry struct {
	Symbol         string `json:"symbol"`
	InstrumentType string `json:"instrument_type"`
}

type Watchlist struct {
	Name             string           `json:"name"`
	WatchlistEntries []WatchlistEntry `json:"watchlist-entries"`
}

type Watchlists struct {
	Items []Watchlist `json:"items"`
}

type WatchlistsResponse struct {
	Context string     `json:"context"`
	Data    Watchlists `json:"data"`
}

type WatchlistResponse struct {
	Context string    `json:"context"`
	Data    Watchlist `json:"data"`
}

type PostWatchlists struct {
	Name             string           `json:"name"`
	WatchlistEntries []WatchlistEntry `json:"watchlist-entries"`
}

type PutWatchlists PostWatchlists

var publicWatchlistsUrl = fmt.Sprintf("%s/public-watchlists", baseUrl)

var privateWatchlistsUrl = fmt.Sprintf("%s/watchlists", baseUrl)

func (t TTClient) getWatchlists(watchlistsUrl string) (WatchlistsResponse, error) {
	resp, err := t.get(watchlistsUrl)

	returnError := func(err error) (WatchlistsResponse, error) {
		return WatchlistsResponse{}, err
	}

	if err != nil {
		return returnError(err)
	}

	responseBody := WatchlistsResponse{}

	err = convertResponseToJson(*resp, &responseBody)

	if err != nil {
		return returnError(err)
	}

	fmt.Println("What are the results?", responseBody)

	return responseBody, nil
}

func (t TTClient) getWatchlist(watchlistUrl, listName string) (WatchlistResponse, error) {
	url := fmt.Sprintf("%s/%s", watchlistUrl, url.PathEscape(listName))

	returnError := func(err error) (WatchlistResponse, error) {
		return WatchlistResponse{}, err
	}

	resp, err := t.get(url)

	if err != nil {
		return returnError(err)
	}

	responseBody := WatchlistResponse{}

	err = convertResponseToJson(*resp, &responseBody)

	if err != nil {
		return returnError(err)
	}

	return responseBody, nil
}

/*
Get all of TastyWorks public watchlists
*/
func (t TTClient) GetPublicWatchlists() (WatchlistsResponse, error) {
	return t.getWatchlists(publicWatchlistsUrl)
}

/*
Get public watchlist with a specific watchlist name will return a list of symbols on that list.
The information here that will be used for the first iteration of this project is `implied-volatility-index-rank`.
Anything greater than 0.60 should be added to a list for further inspection
*/
func (t TTClient) GetPublicWatchlist(listName, instrumentType string) (WatchlistResponse, error) {
	responseBody, err := t.getWatchlist(publicWatchlistsUrl, listName)

	if err != nil {
		return WatchlistResponse{}, err
	}

	matchedInstrumentEntries := []WatchlistEntry{}

	for _, watchlistEntry := range responseBody.Data.WatchlistEntries {
		if watchlistEntry.InstrumentType == instrumentType {
			matchedInstrumentEntries = append(matchedInstrumentEntries, watchlistEntry)
		}
	}

	responseBody.Data.WatchlistEntries = matchedInstrumentEntries

	return responseBody, nil
}

func (t TTClient) CreatePrivateWatchlist(listName string, watchlistEntries []WatchlistEntry) (WatchlistResponse, error) {
	createWatchlistRequest := PostWatchlists{Name: listName, WatchlistEntries: watchlistEntries}

	requestBody, err := prepareRequestBody(createWatchlistRequest)

	returnError := func(err error) (WatchlistResponse, error) {
		return WatchlistResponse{}, err
	}

	if err != nil {
		return returnError(err)
	}

	resp, err := t.post(privateWatchlistsUrl, &requestBody, true)

	if err != nil {
		return returnError(err)
	}

	responseBody := WatchlistResponse{}

	err = convertResponseToJson(*resp, &responseBody)

	if err != nil {
		return returnError(err)
	}

	return responseBody, nil
}

func (t TTClient) GetPrivateWatchlists() (WatchlistsResponse, error) {
	return t.getWatchlists(privateWatchlistsUrl)
}

func (t TTClient) GetPrivateWatchlist(listName string) (WatchlistResponse, error) {
	return t.getWatchlist(privateWatchlistsUrl, listName)
}

func (t TTClient) DeletePrivateWatchlist(listName string) error {
	url := fmt.Sprintf("%s/%s", privateWatchlistsUrl, url.PathEscape(listName))

	_, err := t.delete(url)

	return err
}

func (t TTClient) UpdatePrivateWatchlist(listName string, watchlistEntries []WatchlistEntry) (WatchlistResponse, error) {
	url := fmt.Sprintf("%s/%s", privateWatchlistsUrl, url.PathEscape(listName))

	updateWatchlistRequest := PutWatchlists{Name: listName, WatchlistEntries: watchlistEntries}

	requestBody, err := prepareRequestBody(updateWatchlistRequest)

	returnError := func(err error) (WatchlistResponse, error) {
		return WatchlistResponse{}, err
	}

	if err != nil {
		return returnError(err)
	}

	resp, err := t.put(url, &requestBody)

	if err != nil {
		return returnError(err)
	}

	responseBody := WatchlistResponse{}

	err = convertResponseToJson(*resp, &responseBody)

	if err != nil {
		return returnError(err)
	}

	return responseBody, nil
}
