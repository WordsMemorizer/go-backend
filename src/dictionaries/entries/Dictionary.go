package entries

type Dictionary struct {
	// Id is dictionary's identifier. Can be null if this entities is received from client while It's creating new dictionary
	Id string `json:"id,omitempty"`
	// LanguageToLearn is identifier of the language which user wants to learn. Can be a string in ISO 639-3 or just string (for custom languages or something like this
	LanguageToLearn string `json:"languageToLearn"`
	// LanguageInLearn is identifier of the language which user wants to use to learn. Can be a string in ISO 639-3 or just string (for custom languages or something like this
	LanguageInLearn string `json:"languageInLearn"`
}
