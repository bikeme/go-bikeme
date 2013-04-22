package capitalbikeshare

import (
	"go-bikeme/station"
	"io/ioutil"
	"net/http"
)

const URL string = "http://www.capitalbikeshare.com/data/stations/bikeStations.xml"

func Stations() (stations []station.Station) {
	response, _ := http.Get(URL)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	stations = Parse(body)

	return
}
