package flightserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

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
