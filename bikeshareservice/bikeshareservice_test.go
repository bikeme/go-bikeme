package bikeshareservice

import (
	"fmt"
	"go-bikeme/station"
	"io"
	"net/http"
	"testing"
)

func TestNotInitializedError(t *testing.T) {
	service := BaseService{}
	_, err := service.Stations()
	if err == nil || err.Error() != "IService is not intialized" {
		t.Error("BaseService does not enforce initialization")
	}
}

func TestInitIsAbstract(t *testing.T) {
	service := BaseService{}
	err := service.Init()
	if err == nil || err.Error() != "Method Init() of IService is not implemented" {
		t.Error("Init() method is not abstract")
	}
}

func TestqueryServiceIsAbstract(t *testing.T) {
	service := BaseService{}
	_, err := service.queryService()
	fmt.Println(err.Error())
	if err == nil || err.Error() != "Method queryService() of IService is not implemented for <nil>" {
		t.Error("queryService method is not abstract")
	}
}

func TestparseIsAbstract(t *testing.T) {
	service := BaseService{}
	_, err := service.parse(nil)
	if err == nil || err.Error() != "Method parse(responseBytes []byte) of IService is not implemented for <nil>" {
		t.Error("queryService method is not abstract")
	}
}

func TestStations(t *testing.T) {
	mockService := mockService{}
	mockService.Init()
	_, err := mockService.Stations()
	if err != nil {
		t.Error("Stations returned an error")
	}

	if mockService.queryServiceCallCounter != 1 {
		t.Error(fmt.Sprintf("Expected 1 call to method queryService(), received: %v", mockService.queryServiceCallCounter))
	}

	if mockService.parserCallCounter != 1 {
		t.Error(fmt.Sprintf("Expected 1 call to method parse(responseBytes []byte), received: %v", mockService.parserCallCounter))
	}
}

//The code below is a mock implementation of the IService interface for test purposes
type mockService struct {
	BaseService
	queryServiceCallCounter int
	parserCallCounter       int
}

func (service *mockService) Init() (err error) {
	service.serviceImpl = service
	service.queryServiceCallCounter = 0
	service.parserCallCounter = 0
	return nil
}

func (service *mockService) queryService() (resp *http.Response, err error) {
	service.queryServiceCallCounter++
	mockResponse := &http.Response{}
	mockResponse.Body = &mockResponseBody{}
	return mockResponse, nil
}

func (service *mockService) parse(responseBytes []byte) (stations []station.Station, err error) {
	service.parserCallCounter++
	return []station.Station{}, nil
}

//The mockResponseBody is used to read an array of bytes from the response
type mockResponseBody struct{}

func (response *mockResponseBody) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (response *mockResponseBody) Close() error {
	return nil
}
