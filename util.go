package flightserv

import "strings"

func ParseBody(s string) string {
	// split by new line
	xs := strings.Split(s, "\n")
	ns := make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, "")
	// split by colon
	xs = strings.Split(s, ":")
	ns = make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, ":")
	// split by comma
	xs = strings.Split(s, ",")
	ns = make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, ",")
	// split by opening curly brace
	xs = strings.Split(s, "{")
	ns = make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, "{")
	// split by closing curly brace
	xs = strings.Split(s, "}")
	ns = make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, "}")
	// split by opening square brace
	xs = strings.Split(s, "[")
	ns = make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, "[")
	// split by closing square brace
	xs = strings.Split(s, "]")
	ns = make([]string, 0)
	for _, si := range xs {
		ns = append(ns, strings.TrimSpace(si))
	}
	s = strings.Join(ns, "]")
	return s
}
