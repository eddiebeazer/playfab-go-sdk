package multiplayer_builds

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

type DeleteMultiplayerBuildResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type UpdateMultiplayerBuildRegionsResponse struct {
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

type CreateMultiplayerBuildRequest struct {
	BuildName                   string                  `json:"BuildName"`
	MultiplayerServerCountPerVm int                     `json:"MultiplayerServerCountPerVm"`
	Ports                       []Port                  `json:"Ports"`
	RegionConfigurations        []BuildRegion           `json:"RegionConfigurations"`
	ContainerFlavor             string                  `json:"ContainerFlavor"`
	ContainerImageReference     ContainerImageReference `json:"ContainerImageReference"`
	VmSize                      string                  `json:"VmSize"`
}

type ContainerImageReference struct {
	ImageName string `json:"ImageName"`
	ImageTag  string `json:"Tag"`
}

type Port struct {
	Name     string `json:"Name"`
	Num      string `json:"Num"`
	Protocol string `json:"Protocol"`
}

type BuildRegion struct {
	MaxServers             int                    `json:"MaxServers"`
	DynamicStandbySettings DynamicStandbySettings `json:"DynamicStandbySettings"`
	Region                 string                 `json:"Region"`
	StandbyServers         int                    `json:"StandbyServers"`
}

type DynamicStandbySettings struct {
	IsEnabled bool `json:"IsEnabled"`
}

type ListMultiplayerBuildsRequest struct {
	PageSize int `json:"PageSize"`
}

type ListMultiplayerBuildsResponse struct {
	Data struct {
		BuildSummaries []BuildSummary `json:"BuildSummaries"`
	} `json:"data"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type BuildSummary struct {
	BuildName            string        `json:"BuildName"`
	BuildId              string        `json:"BuildId"`
	RegionConfigurations []BuildRegion `json:"RegionConfigurations"`
}

type UpdateMultiplayerBuildRegionsRequest struct {
	BuildId      string        `json:"BuildId"`
	BuildRegions []BuildRegion `json:"BuildRegions"`
}

type DeleteMultiplayerBuildRequest struct {
	BuildId string `json:"BuildId"`
}

// CreateBuildWithCustomContainer Create a new multiplayer server build
func CreateBuildWithCustomContainer(reqBody CreateMultiplayerBuildRequest) (CreateMultiplayerBuildResponse, error) {
	jsonData, err := json.MarshalIndent(reqBody, "", "\t")
	if err != nil {
		return CreateMultiplayerBuildResponse{Code: 1}, err
	}

	data := client.SendRequest("/MultiplayerServer/CreateBuildWithCustomContainer", jsonData, false, true)

	var response CreateMultiplayerBuildResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// ListMultiplayerBuilds List all current multiplayer builds
func ListMultiplayerBuilds() (ListMultiplayerBuildsResponse, error) {
	req := &ListMultiplayerBuildsRequest{
		PageSize: 50,
	}
	jsonData, err := json.MarshalIndent(*req, "", "\t")
	if err != nil {
		return ListMultiplayerBuildsResponse{Code: 1}, err
	}

	data := client.SendRequest("/MultiplayerServer/ListBuildSummariesV2", jsonData, false, true)

	var response ListMultiplayerBuildsResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// UpdateMultiplayerBuildRegions Update the build region settings for the multiplayer server
func UpdateMultiplayerBuildRegions(req UpdateMultiplayerBuildRegionsRequest) (UpdateMultiplayerBuildRegionsResponse, error) {
	jsonData, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		return UpdateMultiplayerBuildRegionsResponse{Code: 1}, err
	}

	data := client.SendRequest("/MultiplayerServer/UpdateBuildRegions", jsonData, false, true)

	var response UpdateMultiplayerBuildRegionsResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// DeleteMultiplayerBuild Delete's the current multiplayer server from playfab
func DeleteMultiplayerBuild(req DeleteMultiplayerBuildRequest) (DeleteMultiplayerBuildResponse, error) {
	jsonData, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		return DeleteMultiplayerBuildResponse{Code: 1}, err
	}

	data := client.SendRequest("/MultiplayerServer/DeleteBuild", jsonData, false, true)

	var response DeleteMultiplayerBuildResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, nil
	}

	return response, err
}
