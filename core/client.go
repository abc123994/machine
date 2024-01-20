package core

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	tag  int64
	Conn *websocket.Conn
}
type Handler func(string) int32
type Common struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}
type OUT struct {
	Data string `json:"data"`
}

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
	tag := time.Now().UnixNano()
	client := &Client{tag: tag, Conn: c}

	functionMap := map[string]Handler{
		"1": client.square,
		"2": client.double,
	}

	defer c.Close()
	defer onClose(client)
	onConnect(client)
	for {
		_, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}
		var common Common
		err = json.Unmarshal(message, &common)
		if err != nil {
			log.Println("COMMON PARSE ERROR")
			break
		}
		go functionMap[common.Type](common.Type)
	}

}

func (c *Client) square(string) int32 {
	log.Println("square")
	data := OUT{
		Data: "hello square",
	}
	out, _ := json.Marshal(data)
	c.Send("square", out)
	return 0
}
func (c *Client) double(string) int32 {
	log.Println("double")
	data := OUT{
		Data: "hello double",
	}
	out, _ := json.Marshal(data)
	c.Send("double", out)
	return 0
}
func (c *Client) Send(_type string, data []byte) {
	common := Common{
		Type: _type,
		Data: data,
	}

	send, _ := json.Marshal(common)
	c.Conn.WriteMessage(websocket.TextMessage, send)
}
