package main

import (
  "go-bikeme/bicing"
  "fmt"
)

func main() {
  json     := bicing.Stations()
  stations := bicing.Parse(json)

  fmt.Println(stations)
}
