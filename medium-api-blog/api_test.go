package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const successResponse = `{
  "id": "1234567890",
  "name": "Your Name Here"
}`

const invalidResponse = `{
  "error": {
    "message": "Invalid OAuth access token - Cannot parse access token",
    "type": "OAuthException",
    "code": 190,
    "fbtrace_id": "ACWDFEWRSDF92837"
  }
}`

func handleInvalidResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(invalidResponse))
}

func handleSuccessResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(successResponse))
}

func TestApiService_GetUserDetails_InvalidResponse(t *testing.T) {
	service := NewFacebookService()
	invalidResponseServer := httptest.NewServer(http.HandlerFunc(handleInvalidResponse))
	service.BaseUrl = invalidResponseServer.URL + "/"
	_, apiError, _ := service.GetUserDetails()
	assert.Equal(t, http.StatusUnauthorized, apiError.StatusCode)
	assert.Equal(t, apiError.Error.Message, "Invalid OAuth access token - Cannot parse access token")
}

func TestApiService_GetUserDetails_SuccessResponse(t *testing.T) {
	service := NewFacebookService()
	successResponseServer := httptest.NewServer(http.HandlerFunc(handleSuccessResponse))
	service.BaseUrl = successResponseServer.URL + "/"
	apiResponse, apiError, err := service.GetUserDetails()
	assert.Nil(t, apiError)
	assert.Nil(t, err)
	assert.Equal(t, apiResponse.Name, "Your Name Here")
}
