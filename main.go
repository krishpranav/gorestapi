package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/gorestapi/api/aboutapi"
	"github.com/krishpranav/gorestapi/api/album"
	"log"
	"net/http"
)

var aboutapp = []about.Aboutapi{
	{App: "SuperRestApi", Version: 1, Github: "https://github.com/krishpranav/gorestapi"},
}

var albums = []album.Album{
	{ID: "1", Title: "ExampleOne", Artist: "User1", Price: 39.99},
	{ID: "2", Title: "ExampleTwo", Artist: "User2", Price: 39.99},
	{ID: "3", Title: "ExampleThree", Artist: "User3", Price: 39.99},
	{ID: "4", Title: "ExampleFour", Artist: "User4", Price: 39.99},
	{ID: "5", Title: "ExampleFive", Artist: "User5", Price: 39.99},
}

var mainpage = "Hello World"

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views/dist", true)))

	api := router.Group("/api")
	{
		api.GET("/albums", getAlbums)
		router.GET("/albums/:id", getAlbumByID)
		router.POST("/albums", postAlbums)
		router.GET("/about", getAbout)
		router.GET("/", getMain)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:", err)
	}


	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAbout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, aboutapp)
}

func postAlbums(c *gin.Context) {
	var newAlbum album.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getMain(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, mainpage)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
