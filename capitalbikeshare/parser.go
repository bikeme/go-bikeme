package capitalbikeshare

import (
	"bytes"
	"encoding/xml"
	"go-bikeme/station"
)

type XMLStations struct {
	XMLName     xml.Name     `xml:"stations"`
	XMLStations []XMLStation `xml:"station"`
}

type XMLStation struct {
	XMLName      xml.Name `xml:"station"`
	Id           string   `xml:"id"`
	Name         string   `xml:"name"`
	Lat          float32  `xml:"lat"`
	Long         float32  `xml:"long"`
	NbBikes      int64    `xml:"nbBikes"`
	NbEmptyDocks int64    `xml:"nbEmptyDocks"`
}

func Parse(capitalbikshareXML []byte) (stations []station.Station) {

	xmlStations := &XMLStations{}

	xml.NewDecoder(bytes.NewReader(capitalbikshareXML)).Decode(xmlStations)

	for _, xmlStation := range xmlStations.XMLStations {
		stations = append(stations, createStation(xmlStation))
	}

	return
}

func createStation(xmlStation XMLStation) station.Station {
	stationObject := station.Station{}
	stationObject.StationId = xmlStation.Id
	stationObject.StationName = xmlStation.Name

	stationObject.Status = station.Status{xmlStation.NbBikes, xmlStation.NbEmptyDocks}

	return stationObject
}
