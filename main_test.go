package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"testing"
)

/* perform a test request */
func performRequest(t *testing.T, err error) {
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

/* perform a test file for checking frontend is working */
func performFrontend(t *testing.T) {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./views/dist", true)))

	testapi := r.Group("/test")
	{
		testapi.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}

	r.Run()
}