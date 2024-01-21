package main

import (
	"log"
	"machine_svc/config"
	"machine_svc/core"
	"machine_svc/utils"
	"net/http"
)

func main() {
	config.InitConfig()

	http.HandleFunc("/ws", core.HandleFunc)
	utils.SQLExample()
	utils.RedisExample()
	log.Fatal(http.ListenAndServe(config.Ws_host, nil))
}
