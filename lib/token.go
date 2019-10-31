package wiser

import (
	"golang.org/x/xerrors"
)

// TextToPostingsList ...
func TextToPostingsList(documentID int, text string, gramLength int) (InvertIndexHash, error) {
	//tLen := 0
	//position := 0

	ngram, err := doNgram(text, gramLength)
	if err != nil {
		return nil, xerrors.New("failed to doNgram")
	}
	_ = ngram

	return nil, nil
}

// 端数は N 文字にならない可能性がある
func doNgram(text string, n int) ([]string, error) {
	var ngram []string
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if i+n >= len(runes) {
			ngram = append(ngram, string(runes[i:len(runes)]))
		} else {
			ngram = append(ngram, string(runes[i:i+n]))
		}
	}
	return ngram, nil
}

// TokenToPostingsList hoge
func TokenToPostingsList(documentID int, token string, position int) (InvertIndexValue, error) {

	return InvertIndexValue{
		TokenID: 0,
		PostingsList: []PostingsList{
			{
				DocumentID:     0,
				Positions:      nil,
				PositionsCount: 0,
				Next:           &PostingsList{},
			},
		},
		DocCount:       0,
		PositionsCount: 0,
	}, nil
}
