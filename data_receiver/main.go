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
	Prod  DataProducer
}

func main() {
	fmt.Println("----- Starting Data receiver")
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p          DataProducer
		err        error
		kafkaTopic = "obuData"
	)

	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p)

	return &DataReceiver{
		msgch: make(chan types.ANPRData, 128),
		Prod:  p,
	}, nil
}

func (dr *DataReceiver) ProduceData(data types.ANPRData) error {
	return dr.Prod.ProduceData(data)
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

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected!")
	for {
		var data types.ANPRData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error: ", err)
			continue
		}
		if err := dr.ProduceData(data); err != nil {
			fmt.Println("kafka produce error :", err)
		}
	}
}
