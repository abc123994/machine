package core

import (
	"log"
)

var ws_client map[int64]*Client

func init() {
	log.Println("init core")

	ws_client = make(map[int64]*Client)
}

func onConnect(Client *Client) {

	ws_client[Client.tag] = Client
	log.Println("add a client", Client.tag)
}
func onClose(Client *Client) {
	tag := Client.tag
	delete(ws_client, tag)

	log.Println("delete a client", tag)
}
