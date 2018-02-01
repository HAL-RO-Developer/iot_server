package controller

import (
	"net/http"

	"github.com/HAL-RO-Developer/iot_server/controller/validation"
	"github.com/HAL-RO-Developer/iot_server/model"
	"github.com/gin-gonic/gin"
)

func DeviceReceive(c *gin.Context) {
	// デバイスへの命令検索
	req, ok := validation.SearchMyFunction(c)
	if !ok {
		return
	}

	// デバイスの登録チェック(未登録時エラーを返す)
	ret := model.ExistDeviceByIam(req.DeviceID, "")
	if ret {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "デバイスが登録されていません。",
		})
		return
	}

	// デバイスIDチェック
	res := model.ExistDeviceByIam(req.DeviceID, req.MacAddr)
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "デバイスIDが不正です。",
		})
		return
	}

	// 命令取得
	value := model.GetTaskInfo(req.DeviceID)

	c.JSON(http.StatusOK, value[0].Args[0])
}
