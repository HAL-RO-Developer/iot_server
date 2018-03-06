package controller

import (
	"strings"
	"github.com/trevex/golem"
	"net/http"
	"github.com/HAL-RO-Developer/iot_server/model"
)

type Connection struct {
	DeviceID string
	Conn     golem.Connection
}

var connes []Connection

func GetHandle() func(http.ResponseWriter, *http.Request) {
	return createRouter().Handler()
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.OnConnect(connectHandle)

	return router
}

// connection接続時の処理はここに書く
func connectHandle(conn *golem.Connection, http *http.Request) {
	device_id := strings.Split(http.URL.Path, "/")[3]
	for _, value := range connes {
		if value.DeviceID == device_id {
			connes = removeConnection(connes, Connection{DeviceID: value.DeviceID, Conn: value.Conn })
		}
	}
	connes = append(connes, Connection{DeviceID: device_id, Conn: *conn})
}

func MessageSend(msg model.Message ) {
	for _, value := range connes {
		if value.DeviceID == msg.DeviceID {
			value.Conn.Emit("", msg)
		}
	}
}

// スライスの中身削除
func removeConnection(origin []Connection, search Connection) []Connection {
	result := []Connection{}
	for _, v := range origin {
		if v.DeviceID != search.DeviceID {
			result = append(result, Connection{DeviceID: v.DeviceID, Conn: v.Conn})
		}
	}
	return result
}

//type Msg struct {
//	Msg model.Message `json:"msg"`
//}
