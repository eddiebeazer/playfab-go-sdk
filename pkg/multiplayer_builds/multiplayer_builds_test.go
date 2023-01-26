package multiplayer_builds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFalseEndpoints(t *testing.T) {
	respCreateBuild, _ := CreateBuildWithCustomContainer(CreateMultiplayerBuildRequest{})
	assert.Equal(t, 400, respCreateBuild.Code)
	respListBuilds, _ := ListMultiplayerBuilds()
	assert.Equal(t, 404, respListBuilds.Code)
	respUpdateReg, _ := UpdateMultiplayerBuildRegions(UpdateMultiplayerBuildRegionsRequest{})
	assert.Equal(t, 400, respUpdateReg.Code)
	respRemove, _ := DeleteMultiplayerBuild(DeleteMultiplayerBuildRequest{})
	assert.Equal(t, 400, respRemove.Code)
}
