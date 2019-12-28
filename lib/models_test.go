package wiser

import "testing"

var (
	TestPostingsList1 = PostingsList{
		DocumentID:     1,
		Positions:      []int{1},
		PositionsCount: 1,
	}
	TestPostingsList2 = PostingsList{
		DocumentID:     2,
		Positions:      []int{2},
		PositionsCount: 1,
	}

	TestPostingsLists = []PostingsList{
		TestPostingsList1,
		TestPostingsList2,
	}
)

func TestAppendPostingsList(t *testing.T) {
	// TODO: add validation
	got, _ := AppendPostingsList(TestPostingsLists, 1, 3)
	t.Log(got)
}
