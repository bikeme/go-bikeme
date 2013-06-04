package bikeshareservice

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"go-bikeme/location"
	"go-bikeme/station"
	"appengine"
	"appengine/urlfetch"
	"net/http"
)

const TELOFUN_URL string = "http://www.tel-o-fun.co.il:2470/ExternalWS/Geo.asmx"
const SOAP_QUERY string = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
                              <soapenv:Header/>
                              <soapenv:Body>
                                <tem:GetNearestStations>
                                  <tem:longitude>%s</tem:longitude>
                                  <tem:langitude>%s</tem:langitude>
                                  <tem:radius>%v</tem:radius>
                                  <tem:maxResults>%v</tem:maxResults>
                                </tem:GetNearestStations>
                              </soapenv:Body>
                           </soapenv:Envelope>`
const TEL_AVIV_CENTER_LONGITUDE string = "32.071876"
const TEL_AVIV_CENTER_LATITUDE string = "34.7789"
const RADIOUS int = 15000
const MAX_RESULTS int = 250

type telOFunService struct {
	baseService
}

func NewTelOFunService(context appengine.Context) (*telOFunService) {
	service := telOFunService{}
	service.context = context
	service.serviceImpl = &service
	return &service
}

func (service *telOFunService) queryService() (response *http.Response, err error) {
	soapRequestBody := fmt.Sprintf(SOAP_QUERY, TEL_AVIV_CENTER_LONGITUDE, TEL_AVIV_CENTER_LATITUDE, RADIOUS, MAX_RESULTS)
	client := urlfetch.Client(service.context)
	return client.Post(TELOFUN_URL, "text/xml; charset=\"utf-8\"", bytes.NewBufferString(soapRequestBody))
}

func (service *telOFunService) parse(telofunSoapResponse []byte) (stations []station.Station, err error) {
	envelope := &SoapEnvelope{}
	xml.NewDecoder(bytes.NewReader(telofunSoapResponse)).Decode(envelope)

	if envelope.Body.Fault.Message != "" {
		return nil, errors.New(envelope.Body.Fault.Message)
	}

	for _, soapStation := range envelope.Body.Response.Result.CloseBy.Stations {
		if soapStation.Active == 1 {
			stations = append(stations, service.createStation(soapStation))
		}
	}
	return
}

func (service *telOFunService) createStation(soapStation SoapStation) station.Station {
	stationObject := station.Station{}

	stationObject.StationId = soapStation.Id
	stationObject.StationName = soapStation.EnglishName
	stationObject.Address = location.Address{"", "", "", "", location.NewLocationFromString(soapStation.Latitude, soapStation.Longitude)}
	stationObject.Status = station.Status{soapStation.AvailableBike, soapStation.AvailableDocks}

	return stationObject
}

// Define structs that match the TelOFun Xml response hierarchy
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
