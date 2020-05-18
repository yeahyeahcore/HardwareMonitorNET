package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/HardwareMonitorNET/config"
)

func Init() {

}

func Start() {
	r := gin.Default()
	r.GET("/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/info", func(c *gin.Context) {

	})
	r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
