package entries

import "github.com/WordsMemorizer/go-backend/common/entries"

type PhrasePage struct {
	Value    string                       `json:"values"`
	MetaInfo entries.OffsetPagingMetaInfo `json:"metaInfo"`
}
