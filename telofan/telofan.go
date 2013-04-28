package telofan

import (
	"go-bikeme/station"
	"io/ioutil"
	"net/http"
	"bytes"
  "fmt"
)

const URL string = "http://www.tel-o-fun.co.il:2470/ExternalWS/Geo.asmx"
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


func Stations() (stations []station.Station) {
  soapRequestBody := fmt.Sprintf(SOAP_QUERY, TEL_AVIV_CENTER_LONGITUDE, TEL_AVIV_CENTER_LATITUDE, RADIOUS, MAX_RESULTS)
  httpClient := new(http.Client)
  response, err := httpClient.Post(URL, "text/xml; charset=\"utf-8\"", bytes.NewBufferString(soapRequestBody))
  if err != nil {
      fmt.Printf("An error %s occured\n", err.Error())
      return
  }

  bytesResponse, err := ioutil.ReadAll(response.Body) // probably not efficient, done because the stream isn't always a pure XML stream and I have to fix things (not shown here)
  if err != nil {
      fmt.Printf("An error %s occured\n", err.Error())
      return
  }
  stations = Parse(bytesResponse)

	return
}