package handler

import (
	"inventaris/model"
	repo "inventaris/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	data, err := repo.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func CreateOrder(c *gin.Context) {
	payload := new(model.Order)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = repo.CreateOrder(*payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}
