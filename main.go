package main

import (
  "go-bikeme/bicing"
  "go-bikeme/capitalbikeshare"
  "fmt"
)

func main() {
	json := bicing.Stations()
	stations := parser.Parse(json)

	fmt.Println("There are %d stations in the Bicing system!\n", len(stations))

	stations := capitalbikeshare.Stations()
	fmt.Printf("There are %d stations in the Capital Bikeshare system!\n", len(stations))
}
