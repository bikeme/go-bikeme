package location

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

var (
	bicingResponse = `[{},{"data":"[{\"StationID\":\"1\",\"StationName\":\"01 - C GRAN VIA CORTS CATALANES 760\",\"DisctrictCode\":\"2\",\"AddressGmapsLongitude\":\"2.180042000000000000\",\"AddressGmapsLatitude\":\"41.39795200000000000\",\"StationAvailableBikes\":\"17\",\"StationFreeSlot\":\"5\",\"AddressZipCode\":\"08013\",\"AddressStreet1\":\"Gran Via Corts Catalanes\",\"AddressNumber\":\"760\",\"NearbyStationList\":\"24,369,387,426\",\"StationStatusCode\":\"OPN\"},{\"StationID\":\"2\",\"StationName\":\"02 - PL. TETUAN ,8-9\",\"DisctrictCode\":\"2\",\"AddressGmapsLongitude\":\"2.175169000000000000\",\"AddressGmapsLatitude\":\"41.39427200000000000\",\"StationAvailableBikes\":\"1\",\"StationFreeSlot\":\"26\",\"AddressZipCode\":\"08010\",\"AddressStreet1\":\"Plaza Tetu\u00e1n\",\"AddressNumber\":\"8\",\"NearbyStationList\":\"360,368,387,414\",\"StationStatusCode\":\"OPN\"}]"}]`
)

func Test_DistanceInMeters(t *testing.T) {
	from := Location{32.08755, 34.778938}
	to := Location{32.082702, 34.772544}

	distance := from.DistanceInMeters(&to)
	assert.Equal(t, distance, 808, "DistanceInMeters return an unexpected result")
}
