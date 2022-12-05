package flightserv

import "github.com/google/uuid"

type AircraftType struct {
    UUID uuid.UUID `json:"uuid"`
    Name string `json:"name"`
    Description string `json:"description"`
}

type FlightLog struct {
    UUID uuid.UUID `json:"uuid"`
    AircraftTypeUUID uuid.UUID `json:"aircraftTypeUuid"`
    AircraftType string `json:"aircraftType"`
}