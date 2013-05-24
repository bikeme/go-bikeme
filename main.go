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

		from := location.Location{32.064592, 34.779711}
		startTime := time.Now()
		station.NearestStations(from, stations, 5)
		fmt.Printf("Calculating distance to %v stations took %f seconds\n", len(stations), time.Since(startTime).Seconds())
	}
}
