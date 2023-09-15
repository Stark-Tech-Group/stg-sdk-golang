package domain

import (
	logger "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func parseTransMap(s string) (map[float64]float64, error) {
	m := make(map[float64]float64)
	for _, kv := range strings.Split(s, ",") {
		kvSplit := strings.Split(kv, ":")
		if len(kvSplit) == 2 {
			key, errKey := strconv.ParseFloat(kvSplit[0], 64)
			value, errVal := strconv.ParseFloat(kvSplit[1], 64)

			if errKey != nil {
				logger.Errorf("Error converting key to float: %s", errKey)
				return nil, errKey
			}
			if errVal != nil {
				logger.Errorf("Error converting value to float: %s", errVal)
				return nil, errVal
			}

			m[key] = value
		}
	}
	return m, nil
}

func Transform(telem *TelemetryMessage, route *TransformerValue) (*TelemetryMessage, error) {
	transMap, err := parseTransMap(route.ValueType)

	if err != nil {
		return nil, err
	}

	if val, ok := transMap[route.Value]; ok {
		telem.SetValue(route.PointRef, val)
	} else {
		telem.SetValue(route.PointRef, route.Value)
	}

	return telem, nil
}
