package core

import (
	"log"
	pb_machine "machine_svc/schema/machine"
	"time"

	"google.golang.org/protobuf/proto"
)

var m_ws_client map[string]*Client

func init() {
	log.Println("init core")

	m_ws_client = make(map[string]*Client)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			for _, e := range m_ws_client {
				data := pb_machine.Test{}
				out, _ := proto.Marshal(&data)
				e.Send(string(proto.MessageName(&data)), out)
			}
		}
	}()
}

func onConnect(Client *Client) {

	m_ws_client[Client.tag] = Client
	log.Println("add a client", Client.tag)
}
func onClose(Client *Client) {
	tag := Client.tag
	delete(m_ws_client, tag)

	log.Println("delete a client", tag)
}
