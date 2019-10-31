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
	InitDatabase() error
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
	GetDBPath() string
}

// WiserStoreImpl implementation
type WiserStoreImpl struct {
	DBPath string
}

// errors
var (
	DocumentNotFound = xerrors.New("document not found")
)

// NewWiserStore get new WireStore
func NewWiserStore(dbPath string) (WiserStore, error) {
	wiserStoreImpl := WiserStoreImpl{
		DBPath: dbPath,
	}
	var wiserStore WiserStore
	wiserStore = &wiserStoreImpl
	return wiserStore, nil
}

// InitDatabase initialize DB
func (w *WiserStoreImpl) InitDatabase() error {
	db, err := sql.Open("sqlite3", w.DBPath)
	if err != nil {
		msg := xerrors.Errorf("failed to Open: %w", err)
		log.Println(msg)
		return msg
	}

	// migration queries
	baseDir := "./sql/initialize"
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		filePath := baseDir + "/" + f.Name()
		query, err := ioutil.ReadFile(filePath)
		if err != nil {
			msg := xerrors.Errorf("failed to ReadFile: %w", err)
			log.Println(msg)
			return msg
		}

		_, err = db.Exec(string(query))
		if err != nil {
			msg := xerrors.Errorf("failed to Exec: %w", err)
			log.Println(msg)
			return msg
		}
	}

	// prepared statements

	stmt, err := db.Prepare("INSERT INTO documents (title, body) VALUES (?, ?);")

	if err != nil {
		msg := xerrors.Errorf("failed to Prepare: %w", err)
		log.Println(msg)
		return msg
	}
	_, _ = stmt.Exec("title", "hoge")

	return nil
}

// FinDatabase close DB
// TODO: connection close??
func (w *WiserStoreImpl) FinDatabase() error {
	panic("not implemented")
}

// GetDocumentID get document id
func (w *WiserStoreImpl) GetDocumentID(title string) (int, error) {
	db, err := sql.Open("sqlite3", w.DBPath)
	if err != nil {
		msg := xerrors.Errorf("failed to Open: %w", err)
		log.Println(msg)
		return 0, msg
	}

	query := "SELECT id FROM documents WHERE title = ?;"
	rows, err := db.Query(query, title)
	if err != nil {
		msg := xerrors.Errorf("failed to Query: %w", err)
		log.Println(msg)
		return 0, msg
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, DocumentNotFound
	}

	var documentID int
	if err = rows.Scan(&documentID); err != nil {
		msg := xerrors.Errorf("failed to Scan: %w", err)
		log.Println(msg)
		return 0, msg
	}

	return documentID, nil
}

// GetDocumentTitle get document id
func (w *WiserStoreImpl) GetDocumentTitle(documentID int) (string, error) {
	db, err := sql.Open("sqlite3", w.DBPath)
	if err != nil {
		msg := xerrors.Errorf("failed to Open: %w", err)
		log.Println(msg)
		return "", msg
	}

	query := "SELECT title FROM documents WHERE id = ?;"

	rows, err := db.Query(query, documentID)
	if err != nil {
		msg := xerrors.Errorf("failed to Query: %w", err)
		log.Println(msg)
		return "", msg
	}
	defer rows.Close()

	if !rows.Next() {
		return "", DocumentNotFound
	}

	var title string
	if err = rows.Scan(&title); err != nil {
		msg := xerrors.Errorf("failed to Scan: %w", err)
		log.Println(msg)
		return "", msg
	}

	return title, nil
}

// AddDocument if document does not exist then insert else update
func (w *WiserStoreImpl) AddDocument(title string, body string) error {
	db, err := sql.Open("sqlite3", w.DBPath)
	if err != nil {
		msg := xerrors.Errorf("failed to Open: %w", err)
		log.Println(msg)
		return msg
	}
	defer db.Close()

	insertQuery := "INSERT INTO documents (title, body) VALUES (?, ?);"
	updateQuery := "UPDATE documents set body = ? WHERE id = ?;"

	id, err := w.GetDocumentID(title)
	if err != nil {
		if err.Error() == DocumentNotFound.Error() || id == 0 {
			// TODO: survey about error caused by w.DB.Query(insertQuery, title, body)
			_, err = db.Exec(insertQuery, title, body)
			if err != nil {
				msg := xerrors.Errorf("failed to Exec: %w", err)
				log.Println(msg)
				return msg
			}
			
			return nil
		}
		msg := xerrors.Errorf("failed to GetDocumentID: %w", err)
		log.Println(msg)
		return msg
	}

	_, err = db.Exec(updateQuery, body, id)
	if err != nil {
		msg := xerrors.Errorf("failed to Exec update: %w", err)
		log.Println(msg)
		return msg
	}
	return nil
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

// GetDB get *sql.DB for testing...
// TODO: Make it a better implementation
func (w *WiserStoreImpl) GetDBPath() string {
	return w.DBPath
}
