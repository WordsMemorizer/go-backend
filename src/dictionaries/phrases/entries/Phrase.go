package entries

type Phrase struct {
	Id          string     `json:"id"`
	Type        PhraseType `json:"type"`
	Value       string     `json:"value"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"image_url"`
}
