package domain

import (
	logger "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	errKeyMsg = "Error converting key to float: %s"
	errValMsg = "Error converting value to float: %s"
)

func parseTransMap(s string) (map[float64]float64, error) {
	m := make(map[float64]float64)
	for _, kv := range strings.Split(s, ",") {
		kvSplit := strings.Split(kv, ":")
		if len(kvSplit) == 2 {
			key, errKey := strconv.ParseFloat(kvSplit[0], 64)
			value, errVal := strconv.ParseFloat(kvSplit[1], 64)

			if errKey != nil {
				logger.Errorf(errKeyMsg, errKey)
				return nil, errKey
			}
			if errVal != nil {
				logger.Errorf(errValMsg, errVal)
				return nil, errVal
			}

			m[key] = value
		}
	}
	return m, nil
}

func Transform(telem *TelemetryMessage, transVal TransformerValue) error {
	transMap, err := parseTransMap(transVal.ValueType)

	if err != nil {
		return err
	}

	if val, ok := transMap[transVal.Value]; ok {
		telem.SetValue(transVal.PointRef, val)
	} else {
		telem.SetValue(transVal.PointRef, transVal.Value)
	}

	return nil
}
