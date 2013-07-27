package station

import (
	"github.com/stretchrcom/testify/assert"
	"go-bikeme/go-bikeme/location"
	"testing"
	"time"
)

func Test_NearestStations(t *testing.T) {
	from := location.Location{32.064592, 34.779711, ""}
	stations := []Station{
		Station{"1", "", time.Now(), Status{}, location.Location{32.0618, 34.7844, ""}}, //Distance 540m
		Station{"2", "", time.Now(), Status{}, location.Location{32.0657, 34.7831, ""}}, //Distance 342m
		Station{"3", "", time.Now(), Status{}, location.Location{32.0629, 34.7799, ""}}, //Distance 189m
		Station{"4", "", time.Now(), Status{}, location.Location{32.0685, 34.7783, ""}}, //Distance 454m
		Station{"5", "", time.Now(), Status{}, location.Location{32.0653, 34.7766, ""}}, //Distance 303m
	}

	stations = NearestStations(from, stations, 3)

	assert.Equal(t, len(stations), 3, "NearestStations returned too many stations")
	assert.Equal(t, stations[0].StationId, "3", "NearestStations did not sort stations correctly")
	assert.Equal(t, stations[1].StationId, "5", "NearestStations did not sort stations correctly")
	assert.Equal(t, stations[2].StationId, "2", "NearestStations did not sort stations correctly")
}
