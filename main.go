package main

import (
	"fetch.dev/receipt-processor-challenge-submission/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	router.RegisterRoutes(server)

	server.Run(":8080")
}
