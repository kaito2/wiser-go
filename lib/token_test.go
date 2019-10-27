package wiser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextToPostingsList(t *testing.T) {
	_, _ = TextToPostingsList(123, "01弐3四5六", 4)
}

func TestDoNgram(t *testing.T) {
	var data string = "01弐3四5六"
	result, _ := doNgram(data, 3)
	t.Log(result)
	expected := []string{
		"01弐",
		"1弐3",
		"弐3四",
		"3四5",
		"四5六",
		"5六",
		"六",
	}
	assert.Equal(t, expected, result)
}
