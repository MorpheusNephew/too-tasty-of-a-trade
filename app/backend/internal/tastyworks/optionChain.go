package tastyworks

import "fmt"

var optionChainsUrl = fmt.Sprintf("%s/option-chains", baseUrl)

func (t *TTClient) GetOptionChain(symbol string) error {
	fmt.Println("Option chains", optionChainsUrl, symbol)
	url := fmt.Sprintf("%s/%s", optionChainsUrl, symbol)

	resp, err := t.get(url)

	if err != nil {
		return err
	}

	responseBody := make(map[string]any)

	err = convertResponseToJson(resp, &responseBody)

	fmt.Println("What is happening with this option chain?", responseBody)

	return err
}
