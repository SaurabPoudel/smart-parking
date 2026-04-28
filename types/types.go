package types

import "time"

type EventType string

const (
	EventEntry EventType = "ENTRY"
	EventExit  EventType = "Exit"
)

type ParkingSensorData struct {
	SlotID    string    `json:"slot_id"`
	Occupied  bool      `json:"occupied"`
	Timestamp time.Time `json:"timestamp"`
}

type ANPRData struct {
	CameraID  string    `json:"camera_id"`
	Plate     string    `json:"plate"`
	Event     EventType `json:"event"`
	TimeStamp time.Time `json:"timestamp"`
}
