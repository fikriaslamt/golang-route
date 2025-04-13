package handler

import (
	"inventaris/dto"
	repo "inventaris/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStock(c *gin.Context) {
	productID := c.Param("productID")

	data, err := repo.GetStock(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func UpdateStock(c *gin.Context) {
	productID := c.Param("productID")

	payload := new(dto.UpdateStockRequest)

	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := repo.UpdateStock(productID, *payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock updated", "data": data})
}
