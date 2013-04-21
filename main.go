package main

import (
  "bikeme/bicing"
  "bikeme/bicing/parser"
  "fmt"
)

func main() {
  json     := bicing.Stations()
  stations := parser.Parse(json)

  fmt.Println(stations)
}
