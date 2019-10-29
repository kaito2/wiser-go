package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDatabase(t *testing.T) {
	testDBFilePath := "./test.db"
	ws, err := NewWiserStore()
	assert.Nil(t, err)
	ws.InitDatabase(testDBFilePath)
}
