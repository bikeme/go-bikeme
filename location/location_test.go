package location

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func Test_DistanceInMeters(t *testing.T) {
	from := Location{32.08755, 34.778938}
	to := Location{32.082702, 34.772544}

	distance := from.DistanceInMeters(&to)
	assert.Equal(t, distance, 809, "DistanceInMeters return an unexpected result")
}
