package matchmake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFalseEndpoints(t *testing.T) {
	respMatch, _ := CreateMatchmakingQueue(CreateMatchMakingQueueRequest{})
	assert.Equal(t, 400, respMatch.Code)
	respListQueue, _ := ListMatchmakingQueues()
	assert.Equal(t, 200, respListQueue.Code)
	respRemove, _ := RemoveMatchmakingQueue(RemoveMatchmakingQueueRequest{})
	assert.Equal(t, 200, respRemove.Code)
}
