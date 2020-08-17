package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Artist : Model for artist
type Artist struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique;not null"`
}

// CreateArtistInput : struct for create art post request
type CreateArtistInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// FindArtists : Controller for getting all artists
func FindArtists(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var artists []Artist

	db.Find(&artists)

	c.JSON(http.StatusOK, gin.H{"data": artists})
}

// CreateArtist : controller for creating new artists
func CreateArtist(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateArtistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create artist
	artist := Artist{Name: input.Name, Email: input.Email}
	db.Create(&artist)

	c.JSON(http.StatusOK, gin.H{"data": artist})
}
