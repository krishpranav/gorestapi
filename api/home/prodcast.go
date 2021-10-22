package home

type Prodcasts struct {
	ProdcastAuthor string  `json:"prodcastauthor"`
	Comments       string  `json:"comments"`
	StreamedAt     float64 `json:"streamedat"`
	Liked          float64 `json:"liked"`
	Plays          float64 `json:"plays"`
}
