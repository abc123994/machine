package core

import (
	"log"
	pb_common "machine_svc/schema/common"
	pb_machine "machine_svc/schema/machine"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	tag  string
	Conn *websocket.Conn
}
type Handler func([]byte) int32

var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {
		//origin := r.Header.Get("Origin")
		return true
	},
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := &Client{Conn: c}
	log.Println(string(proto.MessageName(&pb_machine.Login{})))
	functionMap := map[string]Handler{
		string(proto.MessageName(&pb_machine.Login{})): client.onlogin,
		"": func([]byte) int32 {
			log.Println("empty")
			return 0
		},
	}

	defer c.Close()
	defer onClose(client)

	for {
		_, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}
		var _c pb_common.Common

		err = proto.Unmarshal(message, &_c)
		if len(message) > 0 {
			if err != nil {
				log.Println("COMMON PARSE ERROR")
				break

			} else {
				log.Println("IN", _c.Type, _c.Data)
				go functionMap[_c.Type](_c.Data)
			}
		}
	}

}

func (c *Client) onlogin(data []byte) int32 {
	var m pb_machine.Login
	err := proto.Unmarshal(data, &m)
	ack := pb_machine.Login_Ack{Result: 0}
	if err != nil {
		ack.Result = 1
	} else {
		c.tag = m.Guid
		onConnect(c)
		ack.Result = 0
	}
	out, err := proto.Marshal(&ack)
	if err != nil {
		panic(err)
	}
	c.Send(string(proto.MessageName(&ack)), out)
	return 0
}

func (c *Client) Send(_type string, data []byte) {
	_c := pb_common.Common{
		Type: _type,
		Data: data,
	}
	send, err := proto.Marshal(&_c)
	if err == nil {

		c.Conn.WriteMessage(websocket.TextMessage, send)
	}
}
