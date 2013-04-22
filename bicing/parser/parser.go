package parser

import (
  "go-bikeme/station"
  "encoding/json"
  "strconv"
)

func Parse(bicingJSON []byte) (stations []station.Station){

  var stationsJSON []interface{}

  json.Unmarshal(bicingJSON, &stationsJSON)

  for _, stationJSON := range stationsJSON {
    stations = append(stations, createStation(stationJSON))
  }

  return
}

func createStation(stationJSON interface{}) station.Station {
  foo := stationJSON.(map[string]interface{})

  stationObject := station.Station{}

  stationObject.StationId = foo["StationID"].(string)

  stationObject.StationName = foo["StationName"].(string)

  availableBikes, _ := strconv.ParseInt(foo["StationAvailableBikes"].(string), 10, 0)
  availableDocks, _ := strconv.ParseInt(foo["StationFreeSlot"].(string), 10, 0)

  stationObject.Status = station.Status{
    availableBikes,
    availableDocks,
  }

  return stationObject
}


