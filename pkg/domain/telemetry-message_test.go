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

func TestTelemetryMessage_MarshalJSON(t *testing.T) {
	m := NewTelemetryMessage("abc")
	m.Ts = 10000
	m.SetValue("a-value", 1.0)
	m.SetValue("b-value", 1.0)

	b, err := json.Marshal(m)
	assert.Nil(t, err)
	s := string(b)

	assert.Equal(t, "{\"a-value\":1,\"b-value\":1,\"deviceId\":\"abc\",\"ts\":10000}", s)
}
