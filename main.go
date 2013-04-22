package main

import (
	"fmt"
	//"go-bikeme/bicing"
	//"go-bikeme/bicing/parser"
	"go-bikeme/capitalbikeshare"
)

func main() {
	//json := bicing.Stations()
	//stations := parser.Parse(json)

	//fmt.Println(len(stations))

	stations := capitalbikeshare.Stations()
	fmt.Printf("There are %d stations in the Capital Bikeshare system!\n", len(stations))
}
