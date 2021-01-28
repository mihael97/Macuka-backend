package models

import "strings"

type TripDto struct {
	Path  []string `json:"path"`
	Start string   `json:"start"`
	End   string   `json:"end"`
	Car   string   `json:"car"`
}

func (trip TripDto) ConnectPath() string {
	return strings.Join(trip.Path, "-")
}
