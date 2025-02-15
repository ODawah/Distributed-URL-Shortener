package handlers

import (
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/ODawah/Distributed-URL-Shortener/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetURL(ctx *gin.Context) {
	Url := &models.URL{} // ✅ Properly initialize the struct

	Url.ID = ctx.Param("shortID")
	fmt.Println("Short ID received:", Url.ID) // Debugging statement

	if len(Url.ID) < 8 {
		fmt.Println("Short ID too short, rejecting request")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := services.GetURL(Url)
	if err != nil {
		fmt.Println("Error retrieving URL from Redis:", err) // Debugging statement
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	now := time.Now()

	// Format and parse it back to time.Time
	formattedTime, err := time.Parse("02-Jan-2006 15:04:05", now.Format("02-Jan-2006 15:04:05"))
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	Request := models.RequestData{
		ShortID:   Url.ID,
		Timestamp: formattedTime,
		IP:        ctx.ClientIP(),
	}

	go services.LogRequestData(Request)

	fmt.Println("Retrieved URL:", Url) // Debugging statement

	ctx.IndentedJSON(http.StatusOK, Url)
}

func ShortenURL(ctx *gin.Context) {
	var url *models.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := services.Shorten(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, url)
}
