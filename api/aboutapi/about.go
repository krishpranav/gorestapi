package aboutapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type aboutapi struct {
	App string `json:"app"`
	Github string `json:"github"`
	Version float64 `json:"version"`
}

var aboutapp = []aboutapi {
	{App: "SuperRestApi", Version: 1, Github: "https://github.com/krishpranav/gorestapi"},
}

func getAbout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, aboutapp)
}