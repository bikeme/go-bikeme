package bikeshareservice

import (
	"encoding/json"
	"go-bikeme/station"
	"net/http"
	"strconv"
)

const BICING_URL string = "https://www.bicing.cat/es/formmap/getJsonObject"

type bicingService struct {
	baseService
	serviceUrl string
}

func NewBicingService() (*bicingService) {
	service := bicingService{}
	service.serviceImpl = &service
	service.serviceUrl = BICING_URL
	return &service
}

func (service *bicingService) queryService() (response *http.Response, err error) {
	return http.Get(service.serviceUrl)
}

func (service *bicingService) parse(bicingJSON []byte) (stations []station.Station, err error) {
	var parsedJSON []interface{}
	json.Unmarshal(bicingJSON, &parsedJSON)

	stationsString := parsedJSON[1].(map[string]interface{})["data"].(string)

	var stationsJSON []interface{}
	json.Unmarshal([]byte(stationsString), &stationsJSON)

	for _, stationJSON := range stationsJSON {
		stations = append(stations, service.createStation(stationJSON))
	}

	return
}

func (service *bicingService) createStation(stationJSON interface{}) station.Station {
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
