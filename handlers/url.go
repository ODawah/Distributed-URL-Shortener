package handlers

import (
	"errors"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/ODawah/Distributed-URL-Shortener/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetURL(ctx *gin.Context) {
	id := ctx.Param("shortID")

	if len(id) < 8 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("request id is less than 8 character")})
		return
	}

	url := services.GetURL(id)
	if url == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	//now := time.Now()
	//
	//// Format and parse it back to time.Time
	//formattedTime, err := time.Parse("02-Jan-2006 15:04:05", now.Format("02-Jan-2006 15:04:05"))
	//if err != nil {
	//	fmt.Println("Error parsing time:", err)
	//}
	//
	//Request := models.RequestData{
	//	ShortID:   id,
	//	Timestamp: formattedTime,
	//	IP:        ctx.ClientIP(),
	//}
	//
	//go services.LogRequestData(Request)

	ctx.IndentedJSON(http.StatusOK, url)
}

func ShortenURL(ctx *gin.Context) {
	var url *models.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	valid := services.IsValidURL(url.URL)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
		return
	}
	err := services.Shorten(url)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, url)
}
