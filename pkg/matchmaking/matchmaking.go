package matchmake

import (
	"dev.azure.com/thedigitalsages/Utilities/_git/playfab-go-sdk/pkg/client"
	"encoding/json"
)

type CreateMultiplayerBuildResponse struct {
	Data struct {
		BuildId string `json:"BuildId"`
	} `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type CreateMatchMakingQueueResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type RemoveMatchmakingQueueResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ListMatchmakingQueuesResult struct {
	Data struct {
		MatchmakingQueues []MatchmakingQueue `json:"MatchMakingQueues"`
	} `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type CreateMatchMakingQueueRequest struct {
	MatchmakingQueue MatchmakingQueue `json:"MatchmakingQueue"`
}

type MatchmakingQueue struct {
	BuildId                 string              `json:"BuildId"`
	MaxMatchSize            int                 `json:"MaxMatchSize"`
	MinMatchSize            int                 `json:"MinMatchSize"`
	Name                    string              `json:"Name"`
	ServerAllocationEnabled bool                `json:"ServerAllocationEnabled"`
	RegionSelectionRule     RegionSelectionRule `json:"RegionSelectionRule"`
	Teams                   []MatchmakingTeam   `json:"Teams"`
}

type MatchmakingTeam struct {
	Name        string `json:"Name"`
	MinTeamSize int    `json:"MinTeamSize"`
	MaxTeamSize int    `json:"MaxTeamSize"`
}

type RegionSelectionRule struct {
	Name       string  `json:"Name"`
	MaxLatency int     `json:"MaxLatency"`
	Path       string  `json:"Path"`
	Weight     float32 `json:"Weight"`
}

type RemoveMatchmakingQueueRequest struct {
	QueueName string `json:"QueueName"`
}

// CreateMatchmakingQueue Creates a new matchmaking queue
func CreateMatchmakingQueue(reqBody CreateMatchMakingQueueRequest) (CreateMatchMakingQueueResponse, error) {
	jsonData, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		return CreateMatchMakingQueueResponse{Code: 1}, err
	}

	data := client.SendRequest("/Match/SetMatchmakingQueue", jsonData, false, true)
	var response CreateMatchMakingQueueResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// ListMatchmakingQueues List all matchmaking queue's
func ListMatchmakingQueues() (ListMatchmakingQueuesResult, error) {
	data := client.SendRequest("/Match/ListMatchmakingQueues", []byte{}, false, true)
	var response ListMatchmakingQueuesResult
	err := json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// RemoveMatchmakingQueue Remove the matchmaking queue's from Playfab
func RemoveMatchmakingQueue(reqBody RemoveMatchmakingQueueRequest) (RemoveMatchmakingQueueResponse, error) {
	jsonData, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		return RemoveMatchmakingQueueResponse{Code: 1}, err
	}

	data := client.SendRequest("/Match/RemoveMatchmakingQueue", jsonData, false, true)
	var response RemoveMatchmakingQueueResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
