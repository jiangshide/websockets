package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var(
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter,r *http.Request){
	var(
		conn *websocket.Conn
		err error
		msgType int
		data []byte
	)
	//Upgrade:websocket
	if conn,err = upgrader.Upgrade(w,r,nil);err != nil{
		return
	}

	//websocket.Conn
	for{
		//Text,Binary
		if msgType,data,err = conn.ReadMessage();err != nil{
			log.Printf("msgType:%s",msgType)
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage,data);err != nil{
			goto ERR
		}
	}

	ERR:
		conn.Close()
}

func main() {
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)//http://localhost:7777/ws
}


