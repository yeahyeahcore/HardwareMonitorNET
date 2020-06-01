package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/HardwareMonitorNET/storage"
)

func handlePostInfo(c *gin.Context) {
	defer c.Request.Body.Close()

	var data storage.Parameter
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	tx, err := storage.Tx()
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = storage.Parameters.Insert(tx, &data)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	c.Status(http.StatusOK)
}
