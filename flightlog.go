package flightserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/url"
)

// GetFlightLogs gets all the flight logs for a specific user.
func (s *Service) GetFlightLogs(UserUUID uuid.UUID) (FlightLogs, dutil.Error) {
	// set path
	s.serv.URL.Path = "log"
	// set query params
	qs := url.Values{"userUUID": {UserUUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return FlightLogs{}, e
	}

	// response structure
	type Data struct {
		FlightLogs `json:"flightLogs"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return FlightLogs{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return FlightLogs{}, e
	}
	return res.Data.FlightLogs, nil
}
