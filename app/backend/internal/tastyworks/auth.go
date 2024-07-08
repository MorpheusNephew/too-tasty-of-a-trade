package tastyworks

import (
	"fmt"
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

type LoginInfo struct {
	User         LoginInfoUserResponse `json:"user"`
	SessionToken string                `json:"session-token"`
}

type LoginInfoResponse struct {
	Context string    `json:"context"`
	Data    LoginInfo `json:"data"`
}

var sessionsUrl = fmt.Sprintf("%s/sessions", baseUrl)

func (t *TTClient) CreateSession(username, password string) error {
	loginInfoRequest := LoginInfoRequest{username, password, true}

	requestBody, err := prepareRequestBody(loginInfoRequest)

	if err != nil {
		return err
	}

	resp, err := t.post(sessionsUrl, &requestBody, false)

	if err != nil {
		return err
	}

	responseBody := LoginInfoResponse{}

	err = convertResponseToJson(*resp, &responseBody)

	if err != nil {
		return err
	}

	t.SessionToken = responseBody.Data.SessionToken

	return nil
}

func (t *TTClient) RemoveSession() (err error) {
	_, err = t.delete(sessionsUrl)

	if err != nil {
		return err
	}

	return nil
}
