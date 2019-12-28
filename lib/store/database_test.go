package store

import (
	"database/sql"
	"os"
	"testing"

	wiser "github.com/kaito2/wiser-go/lib"
	"github.com/stretchr/testify/assert"
)

func TestInitDatabase(t *testing.T) {
	testDBFilePath := "./TestInitDatabase.db"
	defer os.Remove(testDBFilePath)

	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)
}

func TestGetDocumentID(t *testing.T) {
	testDBFilePath := "./TestGetDocumentID.db"
	defer os.Remove(testDBFilePath)

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
	testDBFilePath := "./TestGetDocumentTitle.db"
	defer os.Remove(testDBFilePath)

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
	testDBFilePath := "./TestAddDocument.db"
	defer os.Remove(testDBFilePath)

	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	testTitle := "test_title"
	testBody1 := "test_body1"
	testBody2 := "test_body2"

	// insert test
	err = ws.AddDocument(testTitle, testBody1)
	assert.Nil(t, err)

	// assertion
	db, err := sql.Open("sqlite3", ws.GetDBPath())
	assert.Nil(t, err)
	query := "SELECT body FROM documents WHERE title = ?;"
	rows, err := db.Query(query, testTitle)
	assert.Nil(t, err)
	hasNext := rows.Next()
	assert.True(t, hasNext)
	var gotBody string
	err = rows.Scan(&gotBody)
	assert.Nil(t, err)
	assert.Equal(t, testBody1, gotBody)
	err = rows.Close()
	assert.Nil(t, err)
	assert.Equal(t, testBody1, gotBody)

	// update test
	err = ws.AddDocument(testTitle, testBody2)
	assert.Nil(t, err)

	// assertion
	rows, err = db.Query(query, testTitle)
	defer rows.Close()
	assert.Nil(t, err)
	hasNext = rows.Next()
	assert.True(t, hasNext)
	var gotBody2 string
	err = rows.Scan(&gotBody2)
	assert.Nil(t, err)
	assert.Equal(t, testBody2, gotBody2)

}

func TestGetTokenID(t *testing.T) {
	testDBFilePath := "./TestGetTokenID.db"
	defer os.Remove(testDBFilePath)

	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	testToken := "test_token"
	gotTokenID, gotDocCount, err := ws.GetTokenID(testToken, true)
	assert.Nil(t, err)
	assert.Equal(t, 1, gotTokenID)
	assert.Equal(t, 0, gotDocCount)
}

func TestGetToken(t *testing.T) {
	testDBFilePath := "./TestGetTokenID.db"
	defer os.Remove(testDBFilePath)

	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	// prepare rows
	// assuming that GetTokenID is insert row correctly
	testToken := "test_token"
	gotTokenID, gotDocCount, err := ws.GetTokenID(testToken, true)
	assert.Nil(t, err)
	assert.Equal(t, 1, gotTokenID)
	assert.Equal(t, 0, gotDocCount)

	gotToken, err := ws.GetToken(1)
	assert.Nil(t, err)
	assert.Equal(t, testToken, gotToken)
}

func TestUpdatePostings(t *testing.T) {
	testDBFilePath := "./TestGetTokenID.db"
	defer os.Remove(testDBFilePath)

	ws, err := NewWiserStore(testDBFilePath)
	assert.Nil(t, err)
	err = ws.InitDatabase()
	assert.Nil(t, err)

	// prepare rows
	// assuming that GetTokenID is insert row correctly
	testToken := "test_token"
	gotTokenID, gotDocCount, err := ws.GetTokenID(testToken, true)
	assert.Nil(t, err)
	assert.Equal(t, 1, gotTokenID)
	assert.Equal(t, 0, gotDocCount)

	updatedPostingsList := wiser.PostingsList{
		DocumentID:     123,
		Positions:      []int{1, 2, 3},
		PositionsCount: 3,
		// Next:           nil,
	}
	err = ws.UpdatePostings(1, 1, updatedPostingsList)
	assert.Nil(t, err)

	gotPostingsList, err := ws.GetPostings(1)
	assert.Nil(t, err)
	assert.Equal(t, updatedPostingsList, gotPostingsList)
}
