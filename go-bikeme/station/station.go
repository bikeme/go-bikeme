package station

import (
	"time"
	"fmt"
	"go-bikeme/location"
	"sort"
)

type Status struct {
	AvailableBikes int64
	AvailableDocs  int64
}

type Station struct {
	StationId   string
	StationName string
	LastUpdate  time.Time
	Status      Status
	Address     location.Address
}

func NearestStations(location location.Location, stations []Station, count int) []Station {
	stationDistances := []stationDistance{}
	for _, station := range stations {
		distance := location.DistanceInMeters(&station.Address.Location)
		stationDistances = append(stationDistances, stationDistance{distance, station})
	}

	sds := &stationDistanceSorter{stationDistances}
	sort.Sort(sds)
	sortedStations := []Station{}
	for i := 0; i < count; i++ {
		sortedStations = append(sortedStations, stationDistances[i].station)
		fmt.Printf("%v meters to station:%s, Location: %v\n", stationDistances[i].distance, stationDistances[i].station.StationName, stationDistances[i].station.Address.Location)
	}
	return sortedStations
}

// The code below allows to sort stations by distance
type stationDistance struct {
	distance int64
	station  Station
}

// Implements sort.Interface.
type stationDistanceSorter struct {
	stationDistances []stationDistance
}

func (s *stationDistanceSorter) Len() int {
	return len(s.stationDistances)
}

func (s *stationDistanceSorter) Swap(i, j int) {
	s.stationDistances[i], s.stationDistances[j] = s.stationDistances[j], s.stationDistances[i]
}

func (s *stationDistanceSorter) Less(i, j int) bool {
	return s.stationDistances[i].distance < s.stationDistances[j].distance
}
