package home

type Recommended struct {
	TopTracks string `json:"TopTracks"`
	Plays float64 `json:"plays"`
	ArtistComposed string `json:"artistcomposed"`
	DateReleased float64 `json:"datereleased"`
}
