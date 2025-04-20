package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"inventaris/model"
	repo "inventaris/repository"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	imageStoragePath = "./images"
	maxImageSize     = 5 * 1024 * 1024
)

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// UploadProductImage
func UploadProductImage(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	randomName, err := generateRandomString(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate random filename"})
		return
	}

	if file.Size > maxImageSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the maximum limit of 5 MB"})
		return
	}

	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	isValidExtension := false
	for _, ext := range allowedExtensions {
		if fileExt == ext {
			isValidExtension = true
			break
		}
	}
	if !isValidExtension {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only PNG, JPG, and JPEG are allowed"})
		return
	}

	if err := os.MkdirAll(imageStoragePath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create storage directory"})
		return
	}

	fileName := fmt.Sprintf("%s%s", randomName, fileExt)
	filePath := filepath.Join(imageStoragePath, fileName)

	err = repo.UpdateProduct(id, model.Product{Filepath: fileName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product with image path"})
		return
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "path": filePath})
}

// DownloadProductImage
func DownloadProductImage(c *gin.Context) {
	id := c.Param("id")

	data, err := repo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if data.Filepath == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	filePath := filepath.Join(imageStoragePath, data.Filepath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(filePath)
}

func GetAllProduct(c *gin.Context) {

	data, err := repo.GetAllProducts()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	data, err := repo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func AddProduct(c *gin.Context) {
	payload := new(model.Product)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = repo.AddProduct(*payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product added"})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	payload := new(model.Product)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = repo.UpdateProduct(id, *payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := repo.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted", "id": id})
}
