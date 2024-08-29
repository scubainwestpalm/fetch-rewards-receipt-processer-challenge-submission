package router

import (
	"fmt"
	"net/http"

	"fetch.dev/receipt-processor-challenge-submission/models"
	"github.com/gin-gonic/gin"
)

func processReceipt(ctx *gin.Context) {
	var Receipt models.Receipt
	err := ctx.ShouldBindBodyWithJSON(&Receipt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse receipt data from JSON body."})
		return
	}
	id, err := Receipt.Process()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process receipt data."})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func getPoints(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id not specified."})
		return
	}
	Receipt, err := models.GetReceiptById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"points": Receipt.Points})
}
