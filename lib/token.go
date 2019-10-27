package wiser

import (
	"golang.org/x/xerrors"
)

/*TextToPostingsList ...
 * 渡された文字列から、postings listを作成。
 * @param[in] env 環境
 * @param[in] document_id ドキュメントID。0の場合は、検索キーワードを対象とする。
 * @param[in] text 入力文字列
 * @param[in] text_len 入力文字列の文字長
 * @param[in] n 何-gramか
 * @param[in,out] postings ミニ転置インデックス。NULLを指すポインタを渡すと新規作成
 * @retval 0 成功
 * @retval -1 失敗
 */
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
				CosumentID:     0,
				Positions:      nil,
				PositionsCount: 0,
				Next:           &PostingsList{},
			},
		},
		DocCount:       0,
		PositionsCount: 0,
	}, nil
}
