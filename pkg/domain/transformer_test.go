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

func mockTransformerValue() *TransformerValue {
	return &TransformerValue{
		ValueType: "1:0,2:1,5:2",
		PointRef:  "aRef",
		Value:     5.0,
	}
}

func TestTransform(t *testing.T) {
	telem, err := Transform(mockTelem(), mockTransformerValue())

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 2.0, telem.GetValue("aRef"))
}

func TestTransform_InvalidMapping(t *testing.T) {
	tVal := mockTransformerValue()
	tVal.ValueType = "a string :o"

	telem, err := Transform(mockTelem(), tVal)

	assert.NotNil(t, err)
	assert.Nil(t, telem)
}

func TestTransform_MisMatchedMapping(t *testing.T) {
	tVal := mockTransformerValue()
	tVal.Value = 328.0
	telem, err := Transform(mockTelem(), tVal)

	assert.Nil(t, err)
	assert.NotNil(t, telem)
	assert.Equal(t, 328.0, telem.GetValue("aRef"))
}
