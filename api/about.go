package api

// import "net/http"

type about struct {
    App string `json:"app"`
    Version float64 `json:"version"`
}