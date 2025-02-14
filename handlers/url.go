package handlers

import (
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetURL(ctx *gin.Context) {
	id := ctx.Param("shortID")
	if len(id) < 8 {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	//TODO: check if it's duplicated

	Request := models.RequestData{
		ShortID:   id,
		Timestamp: time.Now(),
		IP:        ctx.ClientIP(),
	}
	fmt.Println(Request)
	ctx.IndentedJSON(http.StatusOK, id)
}

func ShortenURL(ctx *gin.Context) {
	var url models.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := url.Shorten()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, url)
}
