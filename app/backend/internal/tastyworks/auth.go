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

var sessionsUrl = fmt.Sprintf("%s/sessions", baseUrl)

func (t *TTClient) CreateSession(username, password string) (bool, error) {
	loginInfoRequest := LoginInfoRequest{username, password, true}

	jsonBytes, err := json.Marshal(loginInfoRequest)

	if err != nil {
		return false, err
	}

	body := bytes.NewBuffer(jsonBytes)

	resp, err := t.post(sessionsUrl, body)

	if err != nil {
		return false, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	responseBody := LoginInfoResponse{}

	err = json.Unmarshal(bodyBytes, &responseBody)

	if err != nil {
		return false, err
	}

	t.SessionToken = responseBody.Data.SessionToken

	return true, nil
}

func (t *TTClient) RemoveSession() (bool, error) {
	_, err := t.delete(sessionsUrl)

	if err != nil {
		return false, err
	}

	return true, nil
}
