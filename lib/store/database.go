package store

import (
	"database/sql"
	"io/ioutil"
	"log"

	wiser "github.com/kaito2/wiser-go/lib"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/xerrors"
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

// WiserStoreImpl implementation
type WiserStoreImpl struct{}

// NewWiserStore get new WireStore
func NewWiserStore() (WiserStore, error) {
	wiserStoreImpl := WiserStoreImpl{}
	var wiserStore WiserStore
	wiserStore = &wiserStoreImpl
	return wiserStore, nil
}

// InitDatabase initialize DB
func (w *WiserStoreImpl) InitDatabase(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		msg := xerrors.Errorf("failed to Open: %w", err)
		log.Println(msg)
		return msg
	}
	defer db.Close()

	baseDir := "./sql/initialize"
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filePath := baseDir + "/" + f.Name()
		query, err := ioutil.ReadFile(filePath)
		if err != nil {
			return xerrors.Errorf("failed to ReadFile: %w", err)
		}

		_, err = db.Exec(string(query))
		if err != nil {
			msg := xerrors.Errorf("failed to Exec: %w", err)
			log.Println(msg)
			return msg
		}
	}

	return nil
}

func (w *WiserStoreImpl) FinDatabase() error {
	panic("not implemented")
}

func (w *WiserStoreImpl) GetDocumentID(title string) (int, error) {
	panic("not implemented")
}

func (w *WiserStoreImpl) GetDocumentTitle(documentID int) (string, error) {
	panic("not implemented")
}

func (w *WiserStoreImpl) AddDocument(title string, body string) error {
	panic("not implemented")
}

// TODO: Named return value only here
func (w *WiserStoreImpl) GetTokenID(token string, insert bool) (tokenID int, docCount int, err error) {
	panic("not implemented")
}

func (w *WiserStoreImpl) GetToken(tokenID int) (string, error) {
	panic("not implemented")
}

func (w *WiserStoreImpl) GetPostings(tokenID int) (wiser.PostingsList, error) {
	panic("not implemented")
}

func (w *WiserStoreImpl) UpdatePostings(tokenID int, docsCount int, postingsList wiser.PostingsList) {
	panic("not implemented")
}

// Get setting info from DB
func (w *WiserStoreImpl) GetSetting(key string) (string, error) {
	panic("not implemented")
}

func (w *WiserStoreImpl) ReplaceSettings(key string, value string) error {
	panic("not implemented")
}

func (w *WiserStoreImpl) GetDocumentCount() (int, error) {
	panic("not implemented")
}

// TODO: survey how to implement in golang + sqlite
// start DB transaction
func (w *WiserStoreImpl) Begin() {
	panic("not implemented")
}

func (w *WiserStoreImpl) Commit() {
	panic("not implemented")
}

func (w *WiserStoreImpl) RollBack() {
	panic("not implemented")
}
