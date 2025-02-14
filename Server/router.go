package Server

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	return gin.Default()
}
