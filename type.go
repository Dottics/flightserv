package flightserv

import (
	"github.com/google/uuid"
	"time"
)

type AircraftType struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type AircraftTypes []AircraftType

// EqualTo returns a boolean whether a slice of aircraft types are equal.
// Equal in length and equal in order of the aircraft types in the slice.
func (xa *AircraftTypes) EqualTo(aircraftTypes AircraftTypes) bool {
	if len(*xa) != len(aircraftTypes) {
		return false
	}
	for i, a := range *xa {
		aircraftType := aircraftTypes[i]
		if aircraftType != a {
			return false
		}
	}
	return true
}

type FlightLog struct {
	UUID              uuid.UUID `json:"uuid"`
	UserUUID          uuid.UUID `json:"userUuid"`
	AircraftTypeUUID  uuid.UUID `json:"aircraftTypeUuid"`
	AircraftType      string    `json:"aircraftType"`
	Date              time.Time `json:"date"`
	Registration      string    `json:"registration"`
	PilotInCommand    string    `json:"pilotInCommand"`
	Details           string    `json:"details"`
	InstrumentNavAids string    `json:"instrumentNavAids"`
	InstrumentPlace   string    `json:"instrumentPlace"`
	InstrumentActual  string    `json:"instrumentActual"`
	InstrumentFSTD    float64   `json:"instrumentFSTD"`
	InstructorSE      float64   `json:"instructorSE"`
	InstructorME      float64   `json:"instructorME"`
	InstructorFSTD    float64   `json:"instructorFSTD"`
	FSTD              float64   `json:"fstd"`
	EngineType        string    `json:"engineType"`
	DayType           string    `json:"dayType"`
	Dual              float64   `json:"dual"`
	PIC               float64   `json:"pic"`
	PICUS             float64   `json:"picus"`
	Copilot           float64   `json:"copilot"`
	DayLandings       int       `json:"dayLandings"`
	NightLandings     int       `json:"nightLandings"`
	Remarks           string    `json:"remarks"`
}

type FlightLogs []FlightLog

// EqualTo returns a boolean whether a slice of flight logs are equal.
// Equal in length and equal in order of the flight logs in the slice.
func (xf *FlightLogs) EqualTo(flightLogs FlightLogs) bool {
	if len(*xf) != len(flightLogs) {
		return false
	}
	for i, v := range *xf {
		flightLog := flightLogs[i]
		if flightLog != v {
			return false
		}
	}
	return true
}
