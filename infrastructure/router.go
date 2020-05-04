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

	//Tag
	tagController := controllers.NewTagController(NewSqlHandler())
	r.GET("/tags", func(c *gin.Context) { tagController.Index(c) })

	// Template
	templateController := controllers.NewTemplateController(NewSqlHandler())
	r.GET("/templates", func(c *gin.Context) { templateController.Index(c) })
	r.GET("/templates/:uid", func(c *gin.Context) { templateController.Show(c) })
	r.POST("/templates", func(c *gin.Context) { templateController.Create(c) })

	Router = r
}
