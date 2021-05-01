package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanay85/go-micro-product/pkg/api"
	"github.com/maulanay85/go-micro-product/pkg/config"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		category := v1.Group("category")
		category.POST("/", api.CreateCategory)
		category.GET("", api.GetAll)
		category.GET("/:id", api.FindById)
		category.PUT("/:id", api.UpdateById)
		category.DELETE("/:id", api.DeleteById)
	}

	r.Run(config.GetString("server.address"))
}
