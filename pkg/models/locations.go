package models

import (
	"encoding/json"
	"time"
)

type VisitedLocations struct {
	Id         int       `json:"id"`
	Date       time.Time `json:"dateTimeOfVisitLocationPoint"`
	LocationId int       `json:"locationPointId"`
}

type Location struct {
	Latitude  json.RawMessage `json:"latitude"`
	Longitude json.RawMessage `json:"longitude"`
}
