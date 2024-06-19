package tastyworks

import "fmt"

var optionChainsUrl = fmt.Sprintf("%s/option-chains", baseUrl)

func (t *TTClient) GetOptionChain(symbol string) {
	fmt.Println("Option chains", optionChainsUrl)
}
