package bikeshareservice

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"github.com/stretchrcom/testify/mock"
	"go-bikeme/station"
	"io"
	"net/http"
	"testing"
)

func Test_NotInitializedError(t *testing.T) {
	service := BaseService{}
	_, err := service.Stations()

	assert.Equal(t, err.Error(), "IService is not initialized", "BaseService does not enforce initialization")
}

func Test_InitIsAbstract(t *testing.T) {
	service := BaseService{}
	err := service.Init()

	assert.Equal(t, err.Error(), "Method Init() of IService is not implemented", "Init() method is not abstract")
}

func Test_queryServiceIsAbstract(t *testing.T) {
	service := BaseService{}
	_, err := service.queryService()
	fmt.Println(err.Error())

	assert.Equal(t, err.Error(), "Method queryService() of IService is not implemented for <nil>", "queryService() method is not abstract")
}

func Test_parseIsAbstract(t *testing.T) {
	service := BaseService{}
	_, err := service.parse(nil)

	assert.Equal(t, err.Error(), "Method parse(responseBytes []byte) of IService is not implemented for <nil>", "parse(responseBytes []byte) method is not abstract")
}

func Test_Stations(t *testing.T) {
	service := BaseService{}
	mockServiceImpl := new(MockedService)
	service.serviceImpl = mockServiceImpl //Assign the mock IService implementation to serviceImpl

	//Create an HTTP response mock
	mockReader := new(MockedReader)
	mockReader.On("Read", [512]byte{}).Return(0, io.EOF)
	mockReader.On("Close").Return(nil)

	mockResponse := &http.Response{}
	mockResponse.Body = mockReader

	//Expect calls to queryService() and parse([]station.Station)
	mockServiceImpl.On("queryService").Return(mockResponse, nil)
	mockServiceImpl.On("parse", []byte{}).Return([]station.Station{}, nil)

	service.Stations()

	mockServiceImpl.Mock.AssertExpectations(t)
}

//The code below is a mock implementation of the IService interface for test purposes
type MockedService struct {
	mock.Mock
	BaseService
}

func (m *MockedService) queryService() (response *http.Response, err error) {
	args := m.Mock.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockedService) parse(responseBytes []byte) (stations []station.Station, err error) {
	args := m.Mock.Called(responseBytes)
	return args.Get(0).([]station.Station), args.Error(1)
}

//The mockResponseBody is used to read an array of bytes from the response
type MockedReader struct {
	mock.Mock
}

func (m *MockedReader) Read(p []byte) (n int, err error) {
	args := m.Mock.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *MockedReader) Close() error {
	args := m.Mock.Called()
	return args.Error(0)
}
