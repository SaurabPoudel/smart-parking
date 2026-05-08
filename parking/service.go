package main

import (
	"fmt"
	"time"

	"github.com/SaurabPoudel/smart-parking/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type StatusType string

const (
	ACTIVE    StatusType = "ACTIVE"
	COMPLETED StatusType = "COMPLETED"
)

type ParkingSessionServicer interface {
	ProcessEntry(types.ANPRData) error
	ProcessExit(types.ANPRData) error
	GetActiveSessions() map[string]*types.ParkingSession
	GetSessionFee(plate string) (float64, error)
}

// ParkingSession is imported from types package

type ParkingSessionService struct {
	activeSessions    map[string]*types.ParkingSession
	completedSessions []types.ParkingSession
	hourlyRate        float64
}

func NewParkingSessionService() ParkingSessionServicer {
	return &ParkingSessionService{
		activeSessions:    make(map[string]*types.ParkingSession),
		completedSessions: make([]types.ParkingSession, 0),
		hourlyRate:        2.0,
	}
}

func (ps *ParkingSessionService) ProcessEntry(data types.ANPRData) error {
	fmt.Printf("Processing entry for plate: %s\n", data.Plate)

	session := &types.ParkingSession{
		SessionID: uuid.New().String(),
		Plate:     data.Plate,
		SlotID:    "", // Will be assigned by slot manager
		EntryTime: data.TimeStamp,
		Status:    string(ACTIVE),
	}

	ps.activeSessions[data.Plate] = session
	return nil
}

func (ps *ParkingSessionService) ProcessExit(data types.ANPRData) error {
	fmt.Printf("Processing exit for plate: %s\n", data.Plate)

	session, exists := ps.activeSessions[data.Plate]
	if !exists {
		err := fmt.Errorf("no entry found for plate: %s", data.Plate)
		logrus.Errorf("%v", err)
		return err
	}

	session.ExitTime = data.TimeStamp
	session.Duration = session.ExitTime.Sub(session.EntryTime)
	session.Fee = calculateFee(session.Duration, ps.hourlyRate)
	session.Status = string(COMPLETED)

	ps.completedSessions = append(ps.completedSessions, *session)
	delete(ps.activeSessions, data.Plate)

	return nil
}

func (ps *ParkingSessionService) GetActiveSessions() map[string]*types.ParkingSession {
	return ps.activeSessions
}

func (ps *ParkingSessionService) GetSessionFee(plate string) (float64, error) {
	session, exists := ps.activeSessions[plate]
	if !exists {
		return 0, fmt.Errorf("session not found for plate: %s", plate)
	}
	return session.Fee, nil
}

func calculateFee(duration time.Duration, hourlyRate float64) float64 {
	hours := duration.Hours()
	return hours * hourlyRate
}
