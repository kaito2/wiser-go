package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDatabase(t *testing.T) {
	testDBFilePath := "./test.db"
	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)
}

func TestGetDocumentID(t *testing.T) {
	testDBFilePath := "./test.db"
	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	gotDocumentID, err := ws.GetDocumentID("title")
	assert.Nil(t, err)
	expectedDocumentID := 1
	assert.Equal(t, expectedDocumentID, gotDocumentID)
}

func TestGetDocumentTitle(t *testing.T) {
	testDBFilePath := "./test.db"
	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	got, err := ws.GetDocumentTitle(1)
	assert.Nil(t, err)
	expected := "title"
	assert.Equal(t, expected, got)
}

func TestAddDocument(t *testing.T) {
	testDBFilePath := "./test.db"
	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	testTitle := "test_title"
	testBody1 := "test_body1"
	// testBody2 := "test_body2"
	err = ws.AddDocument(testTitle, testBody1)
	assert.Nil(t, err)

	db := ws.GetDB()
	query := "SELECT body FROM documents WHERE title = ?;"
	rows, err := db.Query(query, testTitle)
	assert.Nil(t, err)
	hasNext := rows.Next()
	assert.True(t, hasNext)
	var gotBody string
	err = rows.Scan(&gotBody)
	assert.Nil(t, err)
	t.Log(gotBody)
}
