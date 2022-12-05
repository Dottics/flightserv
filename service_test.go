package flightserv

import (
	"os"
	"testing"
)

func TestNewService(t *testing.T) {
	// set the env vars NewService should automatically get them
	schemeEnv := "FLIGHT_SERVICE_SCHEME"
	scheme := "https"
	err := os.Setenv(schemeEnv, scheme)
	if err != nil {
		t.Errorf("unexpected error setting %s: %v", schemeEnv, err)
	}
	hostEnv := "FLIGHT_SERVICE_HOST"
	host := "flight.dottics.com"
	err = os.Setenv(hostEnv, host)
	if err != nil {
		t.Errorf("unexpected error setting %s: %v", hostEnv, err)
	}
	token := "my-test-token"
	ms := NewService(token)
	if ms.serv.URL.Scheme != scheme {
		t.Errorf("expected Bank Service to have Scheme %s got %s",
			scheme, ms.serv.URL.Scheme,
		)
	}
	if ms.serv.URL.Host != host {
		t.Errorf("expected Bank Service to have Host %s got %s",
			host, ms.serv.URL.Host,
		)
	}
}
