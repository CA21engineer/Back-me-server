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
	// ToDo クエリで検索する
	r.GET("/images", func(c *gin.Context) { imageController.Index(c) })
	r.GET("/images/:id", func(c *gin.Context) { imageController.Show(c) })
	r.POST("/images", func(c *gin.Context) { imageController.Create(c) })

	//Tag
	tagController := controllers.NewTagController(NewSqlHandler())
	r.GET("/tags", func(c *gin.Context) { tagController.Index(c) })

	Router = r
}
