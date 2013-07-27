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
	radius := 6378.14 //earth radius in Km
	dLat := (to.Latitude - from.Latitude) * (math.Pi / 180.0) //Delta latitude in radian
	dLon := (to.Longitude - from.Longitude) * (math.Pi / 180.0) //Delta longitude in radian

	fromLatRad := from.Latitude * (math.Pi / 180.0) //Latitude in radians
	toLatRad := to.Latitude * (math.Pi / 180.0) //Latitude in radians

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(fromLatRad) * math.Cos(toLatRad)
	a := a1 + a2

	circleDistance := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)) // Great circle distance in radians

	return int64(radius * circleDistance * 1000) //Multiply by 1000 to get distance in meters
}