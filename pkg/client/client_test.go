package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEntityToken(t *testing.T) {
	response, err := GetEntityToken()
	assert.Nil(t, err)
	assert.Equal(t, response.Code, 200)
}
