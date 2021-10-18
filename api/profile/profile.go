package profile

type Profile struct {
	User string `json:"user"`
	Password string `json:"password"`
	JoinedAt float64 `json:"joined-at"`
	Toptracks string `json:"toptracks"`
	Following string `json:"following"`
	Followers float64 `json:"followers"`
}