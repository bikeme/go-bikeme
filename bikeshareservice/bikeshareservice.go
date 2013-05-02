package bikeshareservice

import (
	"errors"
	"fmt"
	"go-bikeme/station"
	"io/ioutil"
	"net/http"
)

// The IService interface will decalre methods required to fetch stations data from a bike share service
type IService interface {
	Init() (err error)
	Stations() (stations []station.Station, err error)
	queryService() (resp *http.Response, err error)
	parse(responseBytes []byte) (stations []station.Station, err error)
}

// The BaseService abstract type will implement the template for fetching and parsing stations data
type BaseService struct {
	serviceImpl IService
}

func (service *BaseService) Init() (err error) {
	return errors.New("Method Init() of IService is not implemented")
}

func (service *BaseService) Stations() (stations []station.Station, err error) {
	if service.serviceImpl == nil {
		return nil, errors.New("IService is not intialized")
	}
	response, err := service.serviceImpl.queryService()
	if err != nil {
		fmt.Printf("An error: '%s', occured\n", err.Error())
		return
	}

	defer response.Body.Close()
	bytesResponse, err := ioutil.ReadAll(response.Body) // probably not efficient, done because the stream isn't always a pure XML stream and I have to fix things (not shown here)
	if err != nil {
		fmt.Printf("An error: '%s', occured\n", err.Error())
		return
	}

	stations, err = service.serviceImpl.parse(bytesResponse)

	return
}

func (service *BaseService) queryService() (resp *http.Response, err error) {
	return nil, errors.New(fmt.Sprintf("Method queryService() of IService is not implemented for %T", service.serviceImpl))
}

func (service *BaseService) parse(responseBytes []byte) (stations []station.Station, err error) {
	return nil, errors.New(fmt.Sprintf("Method parse(responseBytes []byte) of IService is not implemented for %T", service.serviceImpl))
}
