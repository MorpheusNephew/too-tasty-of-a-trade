package tastyworks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func convertResponseToJson[T any](resp http.Response, jsonBody T) error {
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &jsonBody)
}

func prepareRequestBody[T any](requestData T) (bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(requestData)

	if err != nil {
		return bytes.Buffer{}, err
	}

	body := bytes.NewBuffer(jsonBytes)

	return *body, nil
}
