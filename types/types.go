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

type ParkingSession struct {
	SessionID string
	Plate     string
	SlotID    string
	EntryTime time.Time
	ExitTime  time.Time
	Duration  time.Duration
	Fee       float64
	Status    string // ACTIVE, COMPLETED
}

type Invoice struct {
	InvoiceID     string
	Plate         string
	TotalSessions int
	TotalAmount   float64
	PeriodStart   time.Time
	PeriodEnd     time.Time
	Status        string // PENDING, PAID
}
