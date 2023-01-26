package admin

import (
	"dev.azure.com/thedigitalsages/Utilities/_git/playfab-go-sdk/pkg/client"
	"dev.azure.com/thedigitalsages/Utilities/_git/playfab-go-sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

type SetTitleDataRequest struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type GetTitleDataRequest struct {
	Keys []string `json:"Keys"`
}

type SetTitleDataResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type GetTitleDataResponse struct {
	Data struct {
		Raw       map[string]string `json:"Data"`
		TitleData []SetTitleDataRequest
	} `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// SetTitleData Adds/Updates Title Data is PlayFab
func SetTitleData(reqBody SetTitleDataRequest) (SetTitleDataResponse, error) {
	jsonData, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	data := client.SendRequest("/Admin/SetTitleData", jsonData, true, false)
	var response SetTitleDataResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// GetTitleData Get the request title data from PlayFab
func GetTitleData(reqBody GetTitleDataRequest) (GetTitleDataResponse, error) {
	jsonData, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		panic(err)
	}

	data := client.SendRequest("/Admin/GetTitleData", jsonData, true, false)

	dat, _ := utils.BeatifyJson(string(data))
	fmt.Println(dat)

	var response GetTitleDataResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	for key, value := range response.Data.Raw {
		response.Data.TitleData = append(response.Data.TitleData, SetTitleDataRequest{
			Key: key, Value: value,
		})
	}
	return response, nil
}
