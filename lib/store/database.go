package store

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/kaito2/wiser-go/lib"
)

// WiserStore interface
type WiserStore interface {
	InitDatabase(dbPath string) error
	FinDatabase() error
	GetDocumentID(title string) (int, error)
	GetDocumentTitle(documentID int) (string, error)
	AddDocument(title, body string) error
	// TODO: Named return value only here
	GetTokenID(token string, insert bool) (tokenID, docCount int, err error)
	GetToken(tokenID int) (string, error)
	GetPostings(tokenID int) (wiser.PostingsList, error)
	UpdatePostings(tokenID int, docsCount int, postingsList wiser.PostingsList)
	// Get setting info from DB
	GetSetting(key string) (string, error)
	ReplaceSettings(key, value string) error
	GetDocumentCount() (int, error)
	// TODO: survey how to implement in golang + sqlite
	// start DB transaction
	Begin()
	Commit()
	RollBack()
}
