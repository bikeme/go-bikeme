package main

import (
	"fmt"
	"go-bikeme/bikeshareservice"
)

func main() {
	services := []bikeshareservice.Service{bikeshareservice.NewBicingService(), bikeshareservice.NewCapitalBikeShareService(), bikeshareservice.NewTelOFunService()}
	for _, service := range services {
		stations, err := service.Stations()
		if err != nil {
			fmt.Printf("#main() received an error: '%s'\n", err.Error())
			return
		}
		fmt.Printf("There are %d stations in the %T system!\n", len(stations), service)
	}
}
