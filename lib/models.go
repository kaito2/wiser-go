package wiser

// PostingsList ポスティングスリスト 文書IDのリンクリスト
type PostingsList struct {
	// 文書のID
	DocumentID int `json:"document_id"`
	// 特定文書中の位置情報配列
	Positions []int `json:"postings"`
	// 特定文書中の位置情報の数
	PositionsCount int `json:"postings_count"`
	// 次の PostingsList へのリンク
	Next *PostingsList `json:"next"`
}

// InvertIndexValue 転置インデックスのバリュー
type InvertIndexValue struct {
	// トークンID
	TokenID int `json:"token_id"`
	// トークンを含む PostingsList
	PostingsList []PostingsList `json:"postings_list"`
	// トークンを含む文書数
	DocCount int `json:"doc_count"`
	// 文書内でのトークン出現数
	PositionsCount int `json:"postings_count"`
}

// InvertIndexHash 転置インデックス
type InvertIndexHash map[string]InvertIndexValue
