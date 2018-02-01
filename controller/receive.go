package controller

import (
	"net/http"

	"fmt"

	"github.com/HAL-RO-Developer/iot_server/model"
	"github.com/gin-gonic/gin"
)

func DeviceReceive(c *gin.Context) {
	device_id := c.Param("id")
	fmt.Println(device_id)
	// デバイスへの命令検索
	ok := model.ExistDeviceById(device_id)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "デバイスIDが登録されていません",
		})
		return
	}
	// 命令取得
	value := model.GetTaskInfo(device_id)
	if value == nil {
		c.JSON(http.StatusOK, 0)
	}
	c.JSON(http.StatusOK, value[0].Args[0])
}
