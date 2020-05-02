package infrastructure

import (
	"ca-zoooom/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	// Image
	imageController := controllers.NewImageController(NewSqlHandler())
	r.GET("/images", func(c *gin.Context) { imageController.Index(c) })
	r.GET("/images/:id", func(c *gin.Context) { imageController.Show(c) })
	r.POST("/images", func(c *gin.Context) { imageController.Create(c) })
	r.PUT("/images/:id", func(c *gin.Context) { imageController.Update(c) })
	r.DELETE("/images/:id", func(c *gin.Context) { imageController.Destroy(c) })

	Router = r
}
