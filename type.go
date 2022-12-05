package flightserv

import "github.com/google/uuid"

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
	UUID             uuid.UUID `json:"uuid"`
	AircraftTypeUUID uuid.UUID `json:"aircraftTypeUuid"`
	AircraftType     string    `json:"aircraftType"`
}
