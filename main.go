package main

import (
	"flag"
	"log"
	"machine_svc/core"
	"net/http"
)

var addr = flag.String("host", "0.0.0.0:8081", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", core.HandleFunc)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
