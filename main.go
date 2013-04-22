package main

import (
  "go-bikeme/bicing"
  "go-bikeme/bicing/parser"
  "fmt"
)

func main() {
  json     := bicing.Stations()
  stations := parser.Parse(json)

  fmt.Println(stations)
}
