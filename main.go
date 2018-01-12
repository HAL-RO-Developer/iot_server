package main

import "github.com/HAL-RO-Developer/iot_server/router"

func main() {
	r := router.Getrouter()
	r.Run()
}
