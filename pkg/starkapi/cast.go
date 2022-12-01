package starkapi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func toInt64(val string) (int64, error) {
	response, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("failed to convert string to int64")
	}
	return response, err
}

func toInt32(val string) (int32, error) {
	response, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.New("failed to convert string to int32")
	}
	return int32(response), err
}

func toFloat64(val string) (float64, error) {
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, errors.New("failed to convert string to float64")
	}
	return value, nil
}

func cast(typ, raw string) (interface{}, error) {

	switch typ {
	case "bigint":
		return toInt64(raw)
	case "bigint:array":
		split := strings.Split(raw, ",")
		typedArray := make([]int64, len(split))
		for i, raw := range split {
			value, err := toInt64(raw)
			if err != nil {
				return nil, err
			} else {
				typedArray[i] = value
			}
		}

		return typedArray, nil
	case "int":
		return toInt32(raw)
	case "int:array":
		split := strings.Split(raw, ",")
		typedArray := make([]int32, len(split))
		for i, raw := range split {
			value, err := toInt32(raw)
			if err != nil {
				return nil, err
			} else {
				typedArray[i] = value
			}
		}
		return typedArray, nil
	case "float":
		return toFloat64(raw)
	case "float:array":
		split := strings.Split(raw, ",")
		typedArray := make([]float64, len(split))
		for i, raw := range split {
			value, err := toFloat64(raw)
			if err != nil {
				return nil, err
			} else {
				typedArray[i] = value
			}
		}
		return typedArray, nil
	case "text":
		return raw, nil
	case "text:array":
		split := strings.Split(raw, ",")
		typedArray := make([]string, len(split))
		for i, raw := range split {
			typedArray[i] = raw
		}
		return typedArray, nil
	case "":
		return raw, nil
	default:
		return nil, fmt.Errorf("unknown data type [%s]", typ)
	}
}
