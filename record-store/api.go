package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func handleAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "No matching album was found"})
}

func handleAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func handleCreateAlbum(c *gin.Context) {
	var album Album
	err := c.BindJSON(&album)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid album data", "error": err})
		return
	}
	albums = append(albums, album)
	c.JSON(http.StatusCreated, album)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/albums", handleAlbums)
	r.POST("/albums", handleCreateAlbum)
	r.GET("/albums/:id", handleAlbumById)

	return r
}
