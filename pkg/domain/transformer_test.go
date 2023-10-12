package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	mockVal0               = 0.0
	mockVal1               = 1.0
	mockDisplayMap         = "0:off,1:on"
	mockDisplayMapBrackets = "[0:off,1:on]"
)

func mockTelem() *TelemetryMessage {
	return &TelemetryMessage{values: map[string]float64{
		"aRef": 1.00,
	}}
}

func mockTransformerValue() TransformerValue {
	return TransformerValue{
		ValueType: "1:0,2:1,5:2",
		PointRef:  "aRef",
		Value:     5.0,
	}
}

func TestTransform(t *testing.T) {
	telem := mockTelem()
	err := Transform(telem, mockTransformerValue())

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 2.0, telem.GetValue("aRef"))
}

func TestTransform_InvalidMapping(t *testing.T) {
	telem := mockTelem()
	tVal := mockTransformerValue()
	tVal.ValueType = "a string :o"

	err := Transform(telem, tVal)

	assert.NotNil(t, err)
}

func TestTransform_MisMatchedMapping(t *testing.T) {
	telem := mockTelem()
	tVal := mockTransformerValue()
	tVal.Value = 328.0
	err := Transform(telem, tVal)

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 328.0, telem.GetValue("aRef"))
}

func TestTransformDisplay(t *testing.T) {
	displayVal, err := TransformDisplay(mockVal0, mockDisplayMap)
	assert.Nil(t, err)
	assert.Equal(t, "off", displayVal)

	displayVal, err = TransformDisplay(mockVal1, mockDisplayMap)
	assert.Nil(t, err)
	assert.Equal(t, "on", displayVal)
}

func TestTransformDisplay_WithBrackets(t *testing.T) {
	displayVal, err := TransformDisplay(mockVal0, mockDisplayMapBrackets)
	assert.Nil(t, err)
	assert.Equal(t, "off", displayVal)

	displayVal, err = TransformDisplay(mockVal1, mockDisplayMapBrackets)
	assert.Nil(t, err)
	assert.Equal(t, "on", displayVal)
}

func TestTransformDisplay_BadMapping(t *testing.T) {
	displayVal, err := TransformDisplay(mockVal0, "a string :o")
	assert.NotNil(t, err)
	assert.Equal(t, "", displayVal)
}

func TestTransformDisplay_KeyDoesNotExist(t *testing.T) {
	displayVal, err := TransformDisplay(3.0, mockDisplayMap)
	assert.Nil(t, err)
	assert.Equal(t, "", displayVal)
}
