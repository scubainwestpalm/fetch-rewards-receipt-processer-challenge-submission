package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/receipts/process", processReceipt)
	server.GET("/receipts/:id/points", getPoints)
}
