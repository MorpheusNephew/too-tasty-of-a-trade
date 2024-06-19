package tastyworks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type LoginInfoRequest struct {
	Username   string `json:"login"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember-me"`
}

type LoginInfoUserResponse struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	ExternalId  string `json:"external-id"`
	IsConfirmed bool   `json:"is-confirmed"`
}

type LoginInfoDataResponse struct {
	User         LoginInfoUserResponse `json:"user"`
	SessionToken string                `json:"session-token"`
}

type LoginInfoResponse struct {
	Context string                `json:"context"`
	Data    LoginInfoDataResponse `json:"data"`
}

func (t *TTClient) CreateSession(username, password string) {
	url := fmt.Sprintf("%s/sessions", baseUrl)

	loginInfoRequest := LoginInfoRequest{username, password, true}

	jsonBytes, err := json.Marshal(loginInfoRequest)

	if err != nil {
		panic("Not worth it")
	}

	body := bytes.NewBuffer(jsonBytes)

	resp, err := t.post(url, body)

	if err != nil {
		panic("This isn't the deal my guy")
	}

	bodyBytes, _ := io.ReadAll(resp.Body)

	responseBody := LoginInfoResponse{}

	err = json.Unmarshal(bodyBytes, &responseBody)

	if err != nil {
		panic("Something happened with the marshalling")
	}

	t.SessionToken = responseBody.Data.SessionToken
}

func (t *TTClient) RemoveSession() {
	url := fmt.Sprintf("%s/sessions", baseUrl)

	_, err := t.delete(url)

	if err != nil {
		panic("An issue occurred removing session")
	}
}