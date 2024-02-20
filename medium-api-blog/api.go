package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type FacebookProfileResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FacebookApiService struct {
	BaseUrl     string `json:"base_url"`
	AccessToken string `json:"acess_token"`
}

type FacebookApiError struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceId string `json:"fbtrace_id"`
	} `json:"error"`
	StatusCode int `json:"status_code"`
}

func NewFacebookService() *FacebookApiService {
	return &FacebookApiService{
		BaseUrl:     "https://graph.facebook.com/v13.0",
		AccessToken: "<some_token>",
	}
}

func (service *FacebookApiService) GetUserDetails() (*FacebookProfileResponse, *FacebookApiError, error) {
	apiUrl := service.BaseUrl + "/me?fields=id,name&access_token=" + service.AccessToken

	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		var apiErr FacebookApiError
		err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return nil, nil, err
		}
		apiErr.StatusCode = res.StatusCode
		return nil, &apiErr, err
	}

	var profileResponse FacebookProfileResponse
	err = json.Unmarshal(body, &profileResponse)
	if err != nil {
		return nil, nil, err
	}

	return &profileResponse, nil, nil
}
