package route

import (
	"inventaris/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	products := router.Group("/products")
	products.GET("/", handler.GetAllProduct)
	products.GET("/:id", handler.GetProductByID)
	products.POST("/add", handler.AddProduct)
	products.PUT("/update/:id", handler.UpdateProduct)
	products.DELETE("/delete/:id", handler.DeleteProduct)

	// Upload and download image endpoints
	products.POST("/upload/:id", handler.UploadProductImage)
	products.GET("/image/:id", handler.DownloadProductImage)

	stock := router.Group("/stock")
	stock.GET("/:productID", handler.GetStock)
	stock.PUT("/update/:productID", handler.UpdateStock)

	order := router.Group("/order")
	order.GET("/:id", handler.GetOrder)
	order.POST("/add", handler.CreateOrder)

}
