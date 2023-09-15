package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockTelem() *TelemetryMessage {
	return &TelemetryMessage{values: map[string]float64{
		"aRef": 1.00,
	}}
}

func mockRoute() *Route {
	return &Route{
		ValueType: "1:0,2:1,5:2",
		PointRef:  "aRef",
	}
}

func TestTransform(t *testing.T) {
	telem, err := Transform(mockTelem(), mockRoute(), 5.0)

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 2.0, telem.GetValue("aRef"))
}

func TestTransform_InvalidMapping(t *testing.T) {
	route := mockRoute()
	route.ValueType = "a string :o"

	telem, err := Transform(mockTelem(), route, 5.0)

	assert.NotNil(t, err)
	assert.Nil(t, telem)
}

func TestTransform_MisMatchedMapping(t *testing.T) {

	telem, err := Transform(mockTelem(), mockRoute(), 328.0)

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 328.0, telem.GetValue("aRef"))
}
