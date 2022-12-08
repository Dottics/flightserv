package flightserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
	"time"
)

func buildFlightLog() FlightLog {
	return FlightLog{
		UUID:              uuid.MustParse("b9782d74-dd43-484a-8fb0-f9eee8e8bcbb"),
		UserUUID:          uuid.MustParse("910b1072-b999-4a6e-b592-ef110ac5b0eb"),
		AircraftTypeUUID:  uuid.UUID{},
		AircraftType:      "A320",
		Date:              time.Date(2022, time.April, 20, 15, 35, 0, 0, time.UTC),
		Registration:      "",
		PilotInCommand:    "SELF",
		Details:           "HKG-CPT",
		InstrumentNavAids: "",
		InstrumentPlace:   "",
		InstrumentActual:  "",
		InstrumentFSTD:    1,
		InstructorSE:      0,
		InstructorME:      0,
		InstructorFSTD:    0,
		FSTD:              2,
		EngineType:        "",
		DayType:           "",
		Dual:              0,
		PIC:               1.2,
		PICUS:             0,
		Copilot:           3.3,
		DayLandings:       1,
		NightLandings:     0,
		Remarks:           "",
	}
}

func TestService_GetFlightLogs(t *testing.T) {
	tt := []struct {
		name       string
		exchange   *microtest.Exchange
		flightLogs FlightLogs
		e          dutil.Error
	}{
		{
			name: "401",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 401,
					Body: `{
						"message": "Forbidden",
						"data": {},
						"errors": {
							"permission": ["please ensure you have permission"]
						}
					}`,
				},
			},
			flightLogs: FlightLogs{},
			e: &dutil.Err{
				Status: 401,
				Errors: map[string][]string{
					"permission": {"please ensure you have permission"},
				},
			},
		},
		{
			name: "get flight logs",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message": "flight logs found",
						"data": {
							"flightLogs": [
								{
									"uuid": "5dd41ccd-0f83-4379-b1d6-deeb26c3f281",
									"userUuid": "1d50a79a-1651-4e31-afd6-a2ce7d187a9d"
								},
								{
									"uuid": "07507951-8492-4449-95a6-3999594c5f64",
									"userUuid": "1d50a79a-1651-4e31-afd6-a2ce7d187a9d"
								},
								{
									"uuid": "a6900fa8-04fb-4de2-9a8a-3c174f3f6315",
									"userUuid": "1d50a79a-1651-4e31-afd6-a2ce7d187a9d"
								}
							]
						},
						"errors": {}
					}`,
				},
			},
			flightLogs: FlightLogs{
				{
					UUID:     uuid.MustParse("5dd41ccd-0f83-4379-b1d6-deeb26c3f281"),
					UserUUID: uuid.MustParse("1d50a79a-1651-4e31-afd6-a2ce7d187a9d"),
				},
				{
					UUID:     uuid.MustParse("07507951-8492-4449-95a6-3999594c5f64"),
					UserUUID: uuid.MustParse("1d50a79a-1651-4e31-afd6-a2ce7d187a9d"),
				},
				{
					UUID:     uuid.MustParse("a6900fa8-04fb-4de2-9a8a-3c174f3f6315"),
					UserUUID: uuid.MustParse("1d50a79a-1651-4e31-afd6-a2ce7d187a9d"),
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		userUUID := uuid.MustParse("1d50a79a-1651-4e31-afd6-a2ce7d187a9d")
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)
			xfl, e := s.GetFlightLogs(userUUID)
			// test the error
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test the query params are formatted correctly
			testURLQuery := fmt.Sprintf("userUUID=%s", userUUID.String())
			if tc.exchange.Request.URL.RawQuery != testURLQuery {
				t.Errorf("expected query '%s' got '%s'",
					testURLQuery, tc.exchange.Request.URL.RawQuery,
				)
			}
			// test the flight logs
			if !tc.flightLogs.EqualTo(xfl) {
				t.Errorf("expected flight logs %v got %v", tc.flightLogs, xfl)
			}
		})
	}
}

func TestService_GetFlightLog(t *testing.T) {
	tt := []struct {
		name      string
		exchange  *microtest.Exchange
		flightLog FlightLog
		e         dutil.Error
	}{
		{
			name: "401 Forbidden",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 401,
					Body: `{
						"message": "Forbidden",
						"data": {},
						"errors": {
							"permission": ["Please ensure you have permission"]
						}
					}`,
				},
			},
			flightLog: FlightLog{},
			e: &dutil.Err{
				Status: 401,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "200 success",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message": "flight log found",
						"data": {
							"flightLog": {
								"uuid": "b9782d74-dd43-484a-8fb0-f9eee8e8bcbb",
								"userUuid": "910b1072-b999-4a6e-b592-ef110ac5b0eb",
								"aircraftType": "A320",
								"date": "2022-04-20T15:35:00.000Z",
								"pilotInCommand": "SELF",
								"details": "HKG-CPT",
								"instrumentFSTD": 1,
								"fstd": 2,
								"pic": 1.2,
								"copilot": 3.3,
								"dayLandings": 1
							}
						}
					}`,
				},
			},
			flightLog: buildFlightLog(),
			e:         nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	testUserUUID := uuid.MustParse("71b2717c-f6aa-43f0-baf1-0fec660cfa16")
	testUUID := uuid.MustParse("905bae9d-3348-4f31-bdaf-ee7059050f63")

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			f, e := s.GetFlightLog(testUserUUID, testUUID)
			// test error
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test requests
			testQueryParams := fmt.Sprintf("UUID=%s&userUUID=%s", testUUID.String(), testUserUUID.String())
			reqQueryParams := tc.exchange.Request.URL.RawQuery
			if reqQueryParams != testQueryParams {
				t.Errorf("expected query params %v got %v", testQueryParams, reqQueryParams)
			}
			// test flight log
			if f != tc.flightLog {
				t.Errorf("expected flight log %v got %v", tc.flightLog, f)
			}
		})
	}
}
