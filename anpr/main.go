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

const sendInterval = time.Second * 5

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

func genPlate() string {
	province := provinces[rand.Intn(len(provinces))]
	zone := rand.Intn(99) + 1 //01 - 99
	vehicle := vehicleTypes[rand.Intn(len(vehicleTypes))]
	number := rand.Intn(99999) + 1 // 01 - 99999
	return fmt.Sprintf("%s-%02d-%s-%05d", province, zone, vehicle, number)

}

func genCameraIDS(n int) []string {
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		ids[i] = uuid.NewString()
	}
	return ids
}

func main() {
	cameraIDS := genCameraIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	activePlates := make(map[string]time.Time)

	for {
		for i := 0; i < len(cameraIDS); i++ {
			var event types.EventType
			var plate string

			if len(activePlates) == 0 || rand.Float32() < 0.7 {
				plate = genPlate()
				event = types.EventEntry
				activePlates[plate] = time.Now()
			} else {
				event = types.EventExit
				idx := 0
				for p := range activePlates {
					if idx == 0 {
						plate = p
						break
					}
					idx++
				}
				delete(activePlates, plate)
			}

			data := types.ANPRData{
				CameraID:  cameraIDS[i],
				Plate:     plate,
				Event:     event,
				TimeStamp: time.Now(),
			}

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Sent %s event for plate: %s\n", event, plate)
		}
		time.Sleep(sendInterval)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
