package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/gorestapi/api/aboutapi"
	"github.com/krishpranav/gorestapi/api/album"
	"github.com/krishpranav/gorestapi/api/artist"
	"log"
	"net/http"
)

var aboutapp = []about.Aboutapi{
	{App: "gorestapi", Author: "krishpranav", Github: "https://github.com/krishpranav/gorestapi", Version: 1,},
}

var albums = []album.Album{
	{ID: "1", Title: "ExampleOne", Artist: "ArtistOne", Price: 39.99},
	{ID: "2", Title: "ExampleTwo", Artist: "ArtistTwo", Price: 39.99},
	{ID: "3", Title: "ExampleThree", Artist: "ArtistThree", Price: 39.99},
	{ID: "4", Title: "ExampleFour", Artist: "ArtistFour", Price: 39.99},
	{ID: "5", Title: "ExampleFive", Artist: "ArtistFive", Price: 39.99},
}

var artists = []artist.Artist {
	{Name: "ArtistOne", Location: "IN", JoinedAt: 2017},
	{Name: "ArtistTwo", Location: "USA", JoinedAt: 2018},
	{Name: "ArtistThree", Location: "UK", JoinedAt: 2019},
	{Name: "ArtistFour", Location: "IN", JoinedAt: 2020},
	{Name: "ArtistFive", Location: "USA", JoinedAt: 2021},
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
		router.GET("/artist", getArtistData)
		router.POST("/artist", postArtist)
		router.GET("/", getMain)

	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:")
		log.Panic(err)
	}


	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAbout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, aboutapp)
}

func getArtistData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, artists)
}

func postArtist(c *gin.Context) {
	var newArtist artist.Artist

	if err := c.BindJSON(&newArtist); err != nil {
		return
	}

	artists = append(artists, newArtist)
	c.IndentedJSON(http.StatusCreated, newArtist)
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

