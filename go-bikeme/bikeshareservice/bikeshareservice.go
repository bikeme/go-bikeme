package bikeshareservice

import (
	"errors"
	"fmt"
	"go-bikeme/station"
	"io/ioutil"
	"net/http"
	"appengine"
)

// The Service interface will declare methods required to fetch stations data from a bike share service
type Service interface {
	Stations() (stations []station.Station, err error)
	queryService() (response *http.Response, err error)
	parse(responseBytes []byte) (stations []station.Station, err error)
}

// The BaseService abstract type will implement the template for fetching and parsing stations data
type baseService struct {
	serviceImpl Service
	context appengine.Context
}

func (service *baseService) Stations() (stations []station.Station, err error) {
	if service.serviceImpl == nil {
		return nil, errors.New("Service is not initialized")
	}
	response, err := service.serviceImpl.queryService()
	if err != nil {
		fmt.Printf("An error: '%s', occurred\n", err.Error())
		return
	}

	defer response.Body.Close()
	bytesResponse, err := ioutil.ReadAll(response.Body) // probably not efficient, done because the stream isn't always a pure XML stream and I have to fix things (not shown here)
	if err != nil {
		fmt.Printf("An error: '%s', occurred\n", err.Error())
		return
	}

	stations, err = service.serviceImpl.parse(bytesResponse)

	return
}

func (service *baseService) queryService() (response *http.Response, err error) {
	return nil, errors.New(fmt.Sprintf("Method queryService() of Service is not implemented for %T", service.serviceImpl))
}

func (service *baseService) parse(responseBytes []byte) (stations []station.Station, err error) {
	return nil, errors.New(fmt.Sprintf("Method parse(responseBytes []byte) of Service is not implemented for %T", service.serviceImpl))
}
