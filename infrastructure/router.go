package infrastructure

import (
	"ca-zoooom/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	r := gin.Default()

	// Video
	videoController := controllers.NewVideoController(NewSqlHandler())
	r.GET("/videos", func(c *gin.Context) { videoController.Index(c) })
	r.GET("/videos/:id", func(c *gin.Context) { videoController.Show(c) })
	r.POST("/videos", func(c *gin.Context) { videoController.Create(c) })
	r.PUT("/videos/:id", func(c *gin.Context) { videoController.Update(c) })
	r.DELETE("/videos/:id", func(c *gin.Context) { videoController.Destroy(c) })

	Router = r
}
