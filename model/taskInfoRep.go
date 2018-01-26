package model

import (
	"sync"
)

type TaskInfo struct {
	Err      string     `json:"err"`
	Success  string     `json:"success"`
	Port     []PortTask `json:"port"`
	DeviceID string     `json:"device_id"`
}

type PortTask struct {
	PortNo int      `json:"port_no"`
	Func   uint64   `json:"function"`
	Args   []uint16 `json:"args"`
}

type GetDevice struct {
	Pin     string `json:"pin"`
	MacAddr string `json:"mac"`
}

type Message struct {
	DeviceID string        `json:"device_id"`
	MacAddr  string        `json:"mac"`
	Msg    	 []ReturnValue `json:"port"`
}

type ReturnValue struct {
	PortNo uint16   `json:"port"`
	Value  int16    `json:"value"`
}

var taskInfo = []TaskInfo{}
var message = []Message{}
var portInfoM = new(sync.Mutex)
var returnValueM = new(sync.Mutex)

func SetTaskInfo(device_id string, task []PortTask) {
	portInfoM.Lock()
	//flg := true
	for _, value := range taskInfo {
		if value.DeviceID == device_id {
			//flg = false
			taskInfo = remove(taskInfo, TaskInfo{DeviceID: value.DeviceID, Port: value.Port})
		}
	}
	/*
	if flg {
		taskInfo = append(taskInfo, TaskInfo{DeviceID: device_id, Port: task})
	}
	*/

	taskInfo = append(taskInfo, TaskInfo{DeviceID: device_id, Port: task})
	portInfoM.Unlock()
}

/*
	自身への命令取得
*/
func GetTaskInfo(device_id string) []PortTask {
	for _, value := range taskInfo {
		if value.DeviceID == device_id {
			return value.Port
		}
	}
	return nil
}

// スライスの中身削除
func remove(origin []TaskInfo, search TaskInfo) []TaskInfo{
	result := []TaskInfo{}
	for _, v := range origin {
		if v.DeviceID != search.DeviceID{
			result = append(result, TaskInfo{DeviceID: v.DeviceID, Port: v.Port})
		}
	}
	return result
}