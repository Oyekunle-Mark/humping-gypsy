package gypsy

// Story the story type
type Story map[string]Chapter

// Chapter structure of the chapter
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option structure of option
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
