package station

import (
	"go-bikeme/location"
)

type Status struct {
	AvailableBikes int64
	AvailableDocs  int64
}

type Station struct {
	StationId   string
	StationName string
	Status      Status
	Address     location.Address
}
