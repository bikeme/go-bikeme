package location

import (
	"strconv"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

type Address struct {
	DistrictCode string
	ZipCode      string
	Street       string
	Number       string
	Location     Location
}

func NewLocationFromString(lat string, lng string) Location {
	location := Location{}
	location.Latitude, _ = strconv.ParseFloat(lat, 64)

	location.Longitude, _ = strconv.ParseFloat(lng, 64)

	return location
}
