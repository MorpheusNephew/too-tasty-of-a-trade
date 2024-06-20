package tastyworks

import "fmt"

type OptionChainResponse = map[string]any

var optionChainsUrl = fmt.Sprintf("%s/option-chains", baseUrl)

func (t *TTClient) GetOptionChain(symbol string) (*OptionChainResponse, error) {
	fmt.Println("Option chains", optionChainsUrl, symbol)
	url := fmt.Sprintf("%s/%s", optionChainsUrl, symbol)

	resp, err := t.get(url)

	if err != nil {
		return nil, err
	}

	responseBody := OptionChainResponse{}

	err = convertResponseToJson(resp, &responseBody)

	if err != nil {
		return nil, err
	}

	fmt.Println("What is happening with this option chain?", responseBody)

	return &responseBody, nil
}
