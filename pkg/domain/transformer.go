package domain

import (
	logger "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func getTransMap(s string) (map[interface{}]interface{}, error) {
	m := make(map[interface{}]interface{})
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

func Transform(telem *TelemetryMessage, route *Route, resp MsgResp) (*TelemetryMessage, error) {
	transMap, err := getTransMap(route.ValueType)

	if err != nil {
		return nil, err
	}

	if transMap[resp.Value] != nil {
		val := transMap[resp.Value]
		telem.Values[route.PointRef] = val.(float64)
	} else {
		telem.Values[route.PointRef] = resp.Value
	}

	return telem, nil
}
