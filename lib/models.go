package wiser

// PostingsList ポスティングスリスト 文書IDのリンクリスト
type PostingsList struct {
	// 文書のID
	CosumentID int
	// 特定文書中の位置情報配列
	Positions []int
	// 特定文書中の位置情報の数
	PositionsCount int
	// 次の PostingsList へのリンク
	Next *PostingsList
}

// InvertIndexValue 転置インデックスのバリュー
type InvertIndexValue struct {
	// トークンID
	TokenID int
	// トークンを含む PostingsList
	PostingsList []PostingsList
	// トークンを含む文書数
	DocCount int
	// 文書内でのトークン出現数
	PositionsCount int
}

// InvertIndexHash 転置インデックス
type InvertIndexHash map[string]InvertIndexValue
