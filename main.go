package main

import (
	"fmt"
	"go-bikeme/bicing"
	"go-bikeme/capitalbikeshare"
	"go-bikeme/telofan"
)

func main() {
	json := bicing.Stations()
	stations := bicing.Parse(json)
	fmt.Printf("There are %d stations in the Bicing system!\n", len(stations))

	stations = capitalbikeshare.Stations()
	fmt.Printf("There are %d stations in the Capital Bikeshare system!\n", len(stations))

	stations = telofan.Stations()
	fmt.Printf("There are %d stations in the TelOfan system!\n", len(stations))
}
