package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pandutria/go-restapi-gin/Controllers/productcontroller"
	"github.com/pandutria/go-restapi-gin/models"
)

func main() { 
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run("0.0.0.0:8082")
}
