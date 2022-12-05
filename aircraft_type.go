package flightserv

import "github.com/dottics/dutil"

// GetAircraftTypes gets all the aircraft types from the flight service. It
// returns all the aircraft types available.
func (s *Service) GetAircraftTypes() (AircraftTypes, dutil.Error) {
	// set path
	s.serv.URL.Path = "/aircraft-type"
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return AircraftTypes{}, e
	}

	// response structure
	type Data struct {
		AircraftTypes `json:"aircraftTypes"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string `json:"errors"`
	}{}

	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return AircraftTypes{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return AircraftTypes{}, e
	}
	// return aircraft types on success
	return res.Data.AircraftTypes, nil
}
