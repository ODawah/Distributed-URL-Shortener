package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Limiter(ctx *gin.Context) {
	if Limiter.Allow() {
		ctx.Next()
	} else {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
	}
}
