package starkapi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt64(t *testing.T) {
	val := "123"
	expected := int64(123)

	result, err := toInt64(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToInt64Slice(t *testing.T) {
	val := "1,2,3"
	expected := []int64{1, 2, 3}

	result, err := toInt64Slice(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToInt32(t *testing.T) {
	val := "42"
	expected := int32(42)

	result, err := toInt32(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToInt32Slice(t *testing.T) {
	val := "10,20,30"
	expected := []int32{10, 20, 30}

	result, err := toInt32Slice(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToFloat64(t *testing.T) {
	val := "3.14"
	expected := 3.14

	result, err := toFloat64(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToFloat64Slice(t *testing.T) {
	val := "1.1,2.2,3.3"
	expected := []float64{1.1, 2.2, 3.3}

	result, err := toFloat64Slice(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToStringSlice(t *testing.T) {
	val := "apple,banana,orange"
	expected := []string{"apple", "banana", "orange"}

	result, err := toStringSlice(val)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToBoolean(t *testing.T) {
	tValue := "true"
	intTValue := "1"
	ftValue := "1.0"
	fValue := "false"
	intFValue := "0"
	invalidValue := "invalid"

	result, err := toBoolean(tValue)
	assert.NoError(t, err)
	assert.True(t, result)

	result, err = toBoolean(ftValue)
	assert.Error(t, err)

	result, err = toBoolean(intTValue)
	assert.NoError(t, err)
	assert.True(t, result)

	result, err = toBoolean(fValue)
	assert.NoError(t, err)
	assert.False(t, result)

	result, err = toBoolean(intFValue)
	assert.NoError(t, err)
	assert.False(t, result)

	result, err = toBoolean(invalidValue)
	assert.Error(t, err)
}

func TestCast(t *testing.T) {
	tests := []struct {
		typ       string
		raw       string
		expected  interface{}
		expectErr bool
	}{
		{"bigint", "123", int64(123), false},
		{"bigint:array", "1,2,3", []int64{1, 2, 3}, false},
		{"int", "42", int32(42), false},
		{"int:array", "10,20,30", []int32{10, 20, 30}, false},
		{"float", "3.14", 3.14, false},
		{"float:array", "1.1,2.2,3.3", []float64{1.1, 2.2, 3.3}, false},
		{"text", "hello", "hello", false},
		{"text:array", "apple,banana,orange", []string{"apple", "banana", "orange"}, false},
		{"", "unknown", "unknown", false},
		{"invalid", "invalid", nil, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s_%s", test.typ, test.raw), func(t *testing.T) {
			result, err := cast(test.typ, test.raw)

			if test.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
