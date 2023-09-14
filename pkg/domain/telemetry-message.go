package domain

import (
	"encoding/json"
	"time"
)

const (
	deviceId = "deviceId"
	ts       = "ts"
)

type TelemetryMessage struct {
	Ts       int64
	DeviceId string
	Values   map[string]float64
}

func NewTelemetryMessage(deviceId string) *TelemetryMessage {
	n := new(TelemetryMessage)
	n.Ts = time.Now().Unix()
	n.DeviceId = deviceId
	n.Values = make(map[string]float64)
	return n
}

func (m *TelemetryMessage) SetValue(key string, value float64) {
	if m.Values == nil {
		m.Values = make(map[string]float64)
	}
	m.Values[key] = value
}

func (m *TelemetryMessage) GetValue(key string) float64 {
	return m.Values[key]
}

func (m *TelemetryMessage) MarshalJSON() ([]byte, error) {
	r := make(map[string]interface{})
	r[deviceId] = m.DeviceId
	r[ts] = m.Ts
	for k, v := range m.Values {
		r[k] = v
	}
	return json.Marshal(r)
}
