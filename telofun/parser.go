package telofun

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-bikeme/station"
)

type SoapStation struct {
	XMLName            xml.Name `xml:"Station"`
	Id                 string   `xml:"Station_id,attr"`
	Name               string   `xml:"Station_Name,attr"`
	EnglishName        string   `xml:"Eng_Station_Name,attr"`
	Description        string   `xml:"Description,attr"`
	EnglishDescription string   `xml:"Eng_Address,attr"`
	Latitude           string   `xml:"Latitude,attr"`
	Longitude          string   `xml:"Longitude,attr"`
	AvailableBike      int64    `xml:"NumOfAvailableBikes,attr"`
	AvailableDocks     int64    `xml:"NumOfAvailableDocks,attr"`
	Active             int      `xml:"IsActive,attr"`
}

type SoapStationsCloseBy struct {
	XMLName  xml.Name      `xml:"StationsCloseBy"`
	Stations []SoapStation `xml:"Station"`
}

type SoapStationsResult struct {
	XMLName xml.Name            `xml:"GetNearestStationsResult"`
	CloseBy SoapStationsCloseBy `xml:"StationsCloseBy"`
}

type SoapStationsResonse struct {
	XMLName xml.Name           `xml:"GetNearestStationsResponse"`
	Result  SoapStationsResult `xml:"GetNearestStationsResult"`
}

type SoapFault struct {
	XMLName xml.Name "soap:Fault"
	Code    string   `xml:"faultcode"`
	Message string   `xml:"faultstring"`
	Details string   `xml:"detail"`
}

type SoapBody struct {
	XMLName  xml.Name            "soap:Body"
	Fault    SoapFault           "soap:Fault"
	Response SoapStationsResonse `xml:"GetNearestStationsResponse"`
}

type SoapEnvelope struct {
	XMLName xml.Name "soap:Envelope"
	Body    SoapBody "soap:Body"
}

func Parse(telofunSoapResponse []byte) (stations []station.Station) {
	envelope := &SoapEnvelope{}
	xml.NewDecoder(bytes.NewReader(telofunSoapResponse)).Decode(envelope)

	if envelope.Body.Fault.Message != "" {
		fmt.Printf("Xml Fault:\n %v \n", envelope.Body.Fault.Message)
		return
	}

	for _, soapStation := range envelope.Body.Response.Result.CloseBy.Stations {
		if soapStation.Active == 1 {
			stations = append(stations, buildStation(soapStation))
		}
	}
	return stations
}

func buildStation(soapStation SoapStation) station.Station {
	stationObject := station.Station{}

	stationObject.StationId = soapStation.Id
	stationObject.StationName = soapStation.EnglishName
	stationObject.Position = station.Position{soapStation.Longitude, soapStation.Latitude}
	stationObject.Status = station.Status{soapStation.AvailableBike, soapStation.AvailableDocks}

	return stationObject
}
