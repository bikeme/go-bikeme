package station

import (
	"time"
)

type Position struct {
	Longitude string
	Latitude  string
}

type Address struct {
	DistrictCode string
	ZipCode      string
	Street       string
	Number       string
}

type Status struct {
	AvailableBikes int64
	AvailableDocs  int64
}

type Station struct {
	StationId   string
	StationName string
	LastUpdate  time.Time
	Position    Position
	Status      Status
	Address     Address
}
