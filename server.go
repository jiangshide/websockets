package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"websockets/impl"
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
		wsConn *websocket.Conn
		err error
		data []byte
		conn *impl.Connection
	)
	//Upgrade:websocket
	if wsConn,err = upgrader.Upgrade(w,r,nil);err != nil{
		return
	}

	if conn,err = impl.InitConnection(wsConn);err != nil{
		goto ERR
	}

	go func() {
		var(
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1*time.Second)
		}
	}()

	for{
		if data,err = conn.ReadMessage();err != nil{
			goto ERR
		}
		if err = conn.WriteMessage(data);err != nil{
			goto ERR
		}
	}

	ERR:
		conn.Close()//TODO:关闭连接的操作
}

func main() {
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)//http://localhost:7777/ws
}


