package tastyworks

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func convertResponseToJson[T any](resp http.Response, jsonBody *T) error {
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(jsonBody)

	return err
}

func prepareRequestBody[T any](requestData T) (bytes.Buffer, error) {
	body := new(bytes.Buffer)

	err := json.NewEncoder(body).Encode(requestData)

	if err != nil {
		return bytes.Buffer{}, err
	}

	return *body, nil
}
