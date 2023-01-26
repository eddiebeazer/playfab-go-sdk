package admin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestTitleData(t *testing.T) {
	// Setting title data
	randomizedKeys := [5]string{
		"One", "Two", "Three", "Four", "Five",
	}
	randomizedValues := [5]string{
		"Six", "Seven", "Eight", "Nine", "Ten",
	}
	rand.Seed(time.Now().Unix())

	titleDataToTest := make([]SetTitleDataRequest, 0)
	for i := 0; i < 5; i++ {
		newTitleData := SetTitleDataRequest{Key: randomizedKeys[i], Value: randomizedValues[rand.Intn(len(randomizedValues))]}
		_, err := SetTitleData(newTitleData)
		assert.Nil(t, err)
		titleDataToTest = append(titleDataToTest, newTitleData)
	}

	// Testing that the set data set is the same as what is pulled
	titleKeysToFetch := make([]string, 0)
	for _, titleKey := range titleDataToTest {
		titleKeysToFetch = append(titleKeysToFetch, titleKey.Key)
	}

	newGetTitleReq := GetTitleDataRequest{Keys: titleKeysToFetch}
	response, err := GetTitleData(newGetTitleReq)
	assert.Nil(t, err)
	fmt.Println(response.Data.TitleData)
	assert.Equal(t, len(titleDataToTest), len(response.Data.TitleData))
}
