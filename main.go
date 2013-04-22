package main

import (
	"fmt"
	"go-bikeme/bicing"
	"go-bikeme/capitalbikeshare"
)

func main() {
	json := bicing.Stations()
	stations := bicing.Parse(json)
	fmt.Printf("There are %d stations in the Bicing system!\n", len(stations))

	stations = capitalbikeshare.Stations()
	fmt.Printf("There are %d stations in the Capital Bikeshare system!\n", len(stations))
}
