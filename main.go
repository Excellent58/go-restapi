package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

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

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album

	for i, a := range albums {
		if a.ID == id {
			if err := c.BindJSON(&updatedAlbum); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
				return
			}

			if updatedAlbum.ID != "" {
				albums[i].ID = updatedAlbum.ID
			}

			if updatedAlbum.Title != "" {
				albums[i].Title = updatedAlbum.Title
			}

			if updatedAlbum.Artist != "" {
				albums[i].Artist = updatedAlbum.Artist
			}

			if updatedAlbum.Price != 0 {
				albums[i].Price = updatedAlbum.Price
			}

			c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for index, a := range albums {
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}