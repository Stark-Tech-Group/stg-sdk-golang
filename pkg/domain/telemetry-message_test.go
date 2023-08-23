package domain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTelemetryMessageSetValue(t *testing.T) {
	m := NewTelemetryMessage("abc")
	m.SetValue("a-value", 100.0)
	m.SetValue("b-value", 200.0)

	assert.Equal(t, 100.0, m.GetValue("a-value"))
	assert.Equal(t, 200.0, m.GetValue("b-value"))

	m.SetValue("a-value", 101.0)
	assert.Equal(t, 101.0, m.GetValue("a-value"))
}

func TestNewTelemetryMessage(t *testing.T) {
	deviceID := "testDevice"
	message := NewTelemetryMessage(deviceID)

	assert.NotNil(t, message)
	assert.Equal(t, deviceID, message.DeviceId)
	assert.NotEmpty(t, message.Ts)
	assert.NotNil(t, message.values)
}

func TestTelemetryMessage_SetValue(t *testing.T) {
	message := &TelemetryMessage{}
	key := "temperature"
	value := 25.5

	message.SetValue(key, value)

	assert.Equal(t, value, message.values[key])
}

func TestTelemetryMessage_GetValue(t *testing.T) {
	message := &TelemetryMessage{}
	key := "humidity"
	value := 60.0

	message.SetValue(key, value)
	result := message.GetValue(key)

	assert.Equal(t, value, result)
}

func TestTelemetryMessage_MarshalJSON(t *testing.T) {
	deviceID := "testDevice"
	ts := int64(1629216000) // Use a specific timestamp here
	message := &TelemetryMessage{
		DeviceId: deviceID,
		Ts:       ts,
	}
	message.SetValue("temperature", 25.5)
	message.SetValue("humidity", 60.0)

	expectedJSON := `{"deviceId":"testDevice","ts":1629216000,"humidity":60,"temperature":25.5}`

	resultJSON, err := json.Marshal(message)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJSON, string(resultJSON))
}
