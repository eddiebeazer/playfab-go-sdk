package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type GetEntityTokenResponse struct {
	Data struct {
		EntityToken string `json:"EntityToken"`
	} `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

var baseApiUrl, titleId, secretKey, entityKey string
var client *http.Client

// Init Initializes default variables for the playfab sdk
func Init() {
	titleId = os.Getenv("PLAYFAB_TITLE_ID")
	secretKey = os.Getenv("PLAYFAB_SECRET")
	baseApiUrl = "https://" + titleId + ".playfabapi.com"
	client = &http.Client{Timeout: 30 * time.Second}

	entityTokenResponse, err := GetEntityToken()
	if err != nil {
		panic(err)
	}
	entityKey = entityTokenResponse.Data.EntityToken
}

// GetEntityToken Returns an entity token that's used for some API request
func GetEntityToken() (GetEntityTokenResponse, error) {
	data := SendRequest("/Authentication/GetEntityToken", []byte{}, true, false)

	var response GetEntityTokenResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, nil
	}
	return response, nil
}

// SendRequest executes a post request and returns raw json
func SendRequest(url string, json []byte, addSecretHeader bool, addEntityHeader bool) []byte {
	if client == nil {
		Init()
	}
	req, err := http.NewRequest("POST", baseApiUrl+url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	// adding headers to the request
	req.Header.Add("Content-Type", "application/json")
	if addSecretHeader {
		req.Header.Add("X-SecretKey", secretKey)
	}
	if addEntityHeader {
		req.Header.Add("X-EntityToken", entityKey)
	}

	// executing request and reading the result into a []byte
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return respBody
}
