package location

import (
	"math"
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

// Distance between two coordinates based on ‘haversine’ formula
// Based on:  https://github.com/kellydunn/golang-geo/blob/master/point.go and http://www.movable-type.co.uk/scripts/latlong.html
func (from *Location) DistanceInMeters(to *Location) int64 {
	r := 6371.0 //m
	dLat := (to.Latitude - from.Latitude) * (math.Pi / 180.0)
	dLon := (to.Longitude - from.Longitude) * (math.Pi / 180.0)

	lat1 := from.Latitude * (math.Pi / 180.0)
	lat2 := to.Latitude * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return int64(r * c * 1000)
}
