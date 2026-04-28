package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgch chan types.ANPRData
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.ANPRData, 128),
	}
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New ANPR connected client connected!")
	for {
		var data types.ANPRData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error: ", err)
			continue
		}
		fmt.Printf("Camera[%10s] <plate[%s]> at [%s] is on %s state. \n", data.CameraID, data.Plate, data.TimeStamp.Format("2006-01-02 15:04:05"), data.Event)
		// dr.msgch <- data
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	dr.conn = conn
	go dr.wsReceiveLoop()
}

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}
