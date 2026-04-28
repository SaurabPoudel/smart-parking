package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

const sendInterval = time.Second

var provinces = []string{
	"BA", // Bagmati
	"GA", // Gandaki
	"LU", // Lumbini
	"KO", // Koshi
	"MA", // Madhesh
	"KA", // Karnali
	"SU", // Sudurpashchim
}

var vehicleTypes = []string{
	"PA", // Private
	"PR", // Public
	"GO", // Government
	"TO", // Tourist
}

func genEvent() types.EventType {
	events := []types.EventType{
		types.EventEntry,
		types.EventExit,
	}
	return events[rand.Intn(len(events))]
}
func genCameraIDS(n int) []string {
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		ids[i] = uuid.NewString()
	}
	return ids
}

func genPlate() string {
	province := provinces[rand.Intn(len(provinces))]
	zone := rand.Intn(99) + 1 //01 - 99
	vehicle := vehicleTypes[rand.Intn(len(vehicleTypes))]
	number := rand.Intn(99999) + 1 // 01 - 99999
	return fmt.Sprintf("%s-%02d-%s-%05d", province, zone, vehicle, number)

}

func main() {
	cameraIDS := genCameraIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(cameraIDS); i++ {
			plate := genPlate()
			event := genEvent()
			time := time.Now()
			data := types.ANPRData{
				CameraID:  cameraIDS[i],
				Plate:     plate,
				Event:     event,
				TimeStamp: time,
			}

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal()
			}
		}
		time.Sleep(sendInterval)

	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
