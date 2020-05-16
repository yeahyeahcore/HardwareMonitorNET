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

	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": ob,
		})
	})

	r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
}
