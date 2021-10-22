package main

import (
	"testing"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
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
func performFrontend(t *testing.T, err error) {
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

	if err != nil {
		panic(err)
	}

	r.Run()
}

/* perform post request */
func postRequest(t *testing.T, err error) {
	r := gin.Default()

	posttest := r.Group("/posttest")
	{
		posttest.POST("/post", func(c *gin.Context) {

		})
	}

	if err != nil {
		panic(err)
	}

	r.Run()
}
