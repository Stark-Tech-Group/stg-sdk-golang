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

func mockMsgResp() MsgResp {
	return MsgResp{Value: 5}
}

func TestTransform(t *testing.T) {
	telem, err := Transform(mockTelem(), mockRoute(), mockMsgResp())

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 2.0, telem.values["aRef"])
}

func TestTransform_InvalidMapping(t *testing.T) {
	route := mockRoute()
	route.ValueType = "a string :o"

	telem, err := Transform(mockTelem(), route, mockMsgResp())

	assert.NotNil(t, err)
	assert.Nil(t, telem)
}

func TestTransform_MisMatchedMapping(t *testing.T) {
	resp := mockMsgResp()
	resp.Value = 328.0

	telem, err := Transform(mockTelem(), mockRoute(), resp)

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 328.0, telem.values["aRef"])
}
