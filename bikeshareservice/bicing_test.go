package bikeshareservice

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"go-bikeme/station"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	bicingResponse = `[{},{"data":"[{\"StationID\":\"1\",\"StationName\":\"01 - C GRAN VIA CORTS CATALANES 760\",\"DisctrictCode\":\"2\",\"AddressGmapsLongitude\":\"2.180042000000000000\",\"AddressGmapsLatitude\":\"41.39795200000000000\",\"StationAvailableBikes\":\"17\",\"StationFreeSlot\":\"5\",\"AddressZipCode\":\"08013\",\"AddressStreet1\":\"Gran Via Corts Catalanes\",\"AddressNumber\":\"760\",\"NearbyStationList\":\"24,369,387,426\",\"StationStatusCode\":\"OPN\"},{\"StationID\":\"2\",\"StationName\":\"02 - PL. TETUAN ,8-9\",\"DisctrictCode\":\"2\",\"AddressGmapsLongitude\":\"2.175169000000000000\",\"AddressGmapsLatitude\":\"41.39427200000000000\",\"StationAvailableBikes\":\"1\",\"StationFreeSlot\":\"26\",\"AddressZipCode\":\"08010\",\"AddressStreet1\":\"Plaza Tetu\u00e1n\",\"AddressNumber\":\"8\",\"NearbyStationList\":\"360,368,387,414\",\"StationStatusCode\":\"OPN\"}]"}]`
)

func Test_downloadTweets(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintln(w, bicingResponse)
	}))
	defer testServer.Close()

	bicing := &BicingService{}
	bicing.Init()
	bicing.serviceUrl = testServer.URL //Set the Bicing service url to the test server Url

	stations, err := bicing.Stations()
	
	assert.Nil(t, err)
	assert.Equal(t, len(stations), 2, "Expected 2 stations")
	assert.Equal(t, stations[0].StationId, "1", "Unexpected station Id")
	assert.Equal(t, stations[0].StationName, "01 - C GRAN VIA CORTS CATALANES 760", "Unexpected station name")
	assert.Equal(t, stations[0].Status, station.Status{17, 5}, "Unexpected station status")
}
