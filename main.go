package main

import (
	"fmt"
	"go-bikeme/bikeshareservice"
	"go-bikeme/location"
	"go-bikeme/station"
	"time"
)

func main() {
	services := []bikeshareservice.Service{bikeshareservice.NewBicingService(), bikeshareservice.NewCapitalBikeShareService(), bikeshareservice.NewTelOFunService()}
	//services := []bikeshareservice.Service{bikeshareservice.NewTelOFunService()}
	for _, service := range services {
		stations, err := service.Stations()
		if err != nil {
			fmt.Printf("#main() received an error: '%s'\n", err.Error())
			return
		}
		fmt.Printf("There are %d stations in the %T system!\n", len(stations), service)

		from := location.Location{32.065174, 34.776449}
		calculateDistanceToAllStations(from, stations)
	}
}

func calculateDistanceToAllStations(from location.Location, toStations []station.Station) {
	startTime := time.Now()
	for _, station := range toStations {
		from.DistanceInMeters(&station.Address.Location)
		//fmt.Printf("%v meters meters to station:%s\n", distance, station.StationName)
	}
	fmt.Printf("Calculating distance to %v stations took %f seconds\n", len(toStations), time.Since(startTime).Seconds())
}
