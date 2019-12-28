package wiser

import "golang.org/x/xerrors"

// PostingsList ポスティングスリスト 文書IDのリンクリスト
type PostingsList struct {
	// 文書のID
	DocumentID int `json:"document_id"`
	// 特定文書中の位置情報配列
	Positions []int `json:"postings"`
	// 特定文書中の位置情報の数
	PositionsCount int `json:"postings_count"`
	// 次の PostingsList へのリンク
	// Next *PostingsList `json:"next"`
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
type InvertIndexMap map[string]InvertIndexValue

// TODO: token 自体がないなら key を追加する。
// TODO: その Document に関する PostingsList がないなら追加する。
func (iim InvertIndexMap) AppendPostingsList(documentID int, token string, position int) error {
	beforeInvertIndexMap := iim
	for key, _ := range beforeInvertIndexMap {
		if key == token {
			var err error
			updatedPostingsList, err := AppendPostingsList(iim[key].PostingsList, documentID, position)
			iim[key] = InvertIndexValue{
				TokenID:        iim[key].TokenID,
				PostingsList:   updatedPostingsList,
				DocCount:       iim[key].DocCount,
				PositionsCount: iim[key].PositionsCount + 1,
			}
			if err != nil {
				return xerrors.Errorf("failed to AppendPostingsList: %w", err)
			}
		}
	}
	// TODO: add new InvertIndexValue
	return nil
}

func AppendPostingsList(postingsList []PostingsList, documentID int, position int) ([]PostingsList, error) {
	for i, postings := range postingsList {
		if postings.DocumentID == documentID {
			postingsList[i].Positions = append(postingsList[i].Positions, position)
			postingsList[i].PositionsCount = postingsList[i].PositionsCount + 1
			return postingsList, nil
		}
	}
	postingsList = append(postingsList, PostingsList{
		DocumentID:     documentID,
		Positions:      []int{position},
		PositionsCount: 1,
	})
	return postingsList, nil
}
