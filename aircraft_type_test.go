package flightserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetAircraftTypes(t *testing.T) {
	tt := []struct {
		name          string
		exchange      *microtest.Exchange
		aircraftTypes AircraftTypes
		e             dutil.Error
	}{
		{
			name: "successful request",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message": "aircraft types found",
						"data": {
							"aircraftTypes": [
								{"uuid": "57991734-720e-40fc-8bd9-43bbd64cca0d", "name": "A380", "description": "an airbus 380"},
								{"uuid": "07507951-8492-4449-95a6-3999594c5f64", "name": "XJZ12", "description": ""},
								{"uuid": "905bae9d-3348-4f31-bdaf-ee7059050f63", "name": "F210", "description": ""}
							]
						},
						"errors": {}
					}`,
				},
			},
			aircraftTypes: AircraftTypes{
				AircraftType{
					UUID:        uuid.MustParse("57991734-720e-40fc-8bd9-43bbd64cca0d"),
					Name:        "A380",
					Description: "an airbus 380",
				},
				AircraftType{
					UUID: uuid.MustParse("07507951-8492-4449-95a6-3999594c5f64"),
					Name: "XJZ12",
				},
				AircraftType{
					UUID: uuid.MustParse("905bae9d-3348-4f31-bdaf-ee7059050f63"),
					Name: "F210",
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xat, e := s.GetAircraftTypes()
			if !dutil.ErrorEqual(e, tc.e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if len(xat) != len(tc.aircraftTypes) {
				t.Errorf("expected aircraft types length %d got %d", len(tc.aircraftTypes), len(xat))
			}
			if !tc.aircraftTypes.EqualTo(xat) {
				t.Errorf("expected aircraft types %v got %v", tc.aircraftTypes, xat)
			}
		})
	}
}
