package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

/* perfomr get request */
func getrequest(t *testing.T, err error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	if err != nil {
		panic(err)
	}

}
