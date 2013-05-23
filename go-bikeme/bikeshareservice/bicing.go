package bikeshareservice

import (
	"bytes"
	"encoding/json"
	"go-bikeme/station"
	"net/http"
	"strconv"
	"strings"
	"appengine"
	"appengine/urlfetch"
)

const BICING_URL string = "https://www.bicing.cat/es/formmap/getJsonObject"

type bicingService struct {
	baseService
	serviceUrl string
}

func NewBicingService(context appengine.Context) (*bicingService) {
	service := bicingService{}
	service.context = context
	service.serviceImpl = &service
	service.serviceUrl = BICING_URL
	return &service
}

func (service *bicingService) queryService() (response *http.Response, err error) {
	client := urlfetch.Client(service.context)
	return client.Get(service.serviceUrl)
}

func (service *bicingService) parse(bicingJSON []byte) (stations []station.Station, err error) {

	bicingCommands := []bicingCommand{}

	json.NewDecoder(bytes.NewReader(bicingJSON)).Decode(&bicingCommands)

	bicingStations := []bicingJsonStation{}

	json.NewDecoder(strings.NewReader(bicingCommands[1].Data)).Decode(&bicingStations)

	for _, bicingStation := range bicingStations {
		stations = append(stations, service.createStation(bicingStation))
	}

	return
}

func (service *bicingService) createStation(bicingJsonStation bicingJsonStation) station.Station {

	stationObject := station.Station{}

	stationObject.StationId = bicingJsonStation.StationID
	stationObject.StationName = bicingJsonStation.StationName

	availableBikes, _ := strconv.ParseInt(bicingJsonStation.StationAvailableBikes, 10, 0)
	availableDocks, _ := strconv.ParseInt(bicingJsonStation.StationFreeSlot, 10, 0)

	stationObject.Status = station.Status{
		availableBikes,
		availableDocks,
	}

	return stationObject
}

type bicingJsonStation struct {
	StationID string `json:"StationID`
	StationName string `json:"StationName`
	StationAvailableBikes string `json:"StationAvailableBikes`
	StationFreeSlot string `json:"StationFreeSlot`
}

type bicingCommand struct {
	Data string `json:"data"`
}
