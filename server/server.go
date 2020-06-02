package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/HardwareMonitorNET/config"
	"github.com/yeahyeahcore/HardwareMonitorNET/storage"
)

//Init инициализирует тупа усё
func Init() {
	storage.Init()
}

//Start ну стартуем
func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/config_info", handleConfigInfo)
	r.POST("/post_info", handlePostInfo)

	r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
