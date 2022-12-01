package starkapi

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strconv"
	"strings"
)

const arraySep = ","

func toInt64(val string) (int64, error) {
	response, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("failed to convert string to int64")
	}
	return response, err
}

func toInt64Slice(val string) ([]int64, error) {
	split := strings.Split(val, arraySep)
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
}

func toInt32(val string) (int32, error) {
	response, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.New("failed to convert string to int32")
	}
	return int32(response), err
}
func toInt32Slice(val string) ([]int32, error) {
	split := strings.Split(val, arraySep)
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
}

func toFloat64(val string) (float64, error) {
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, errors.New("failed to convert string to float64")
	}
	return value, nil
}

func toFloat64Slice(val string) ([]float64, error) {
	split := strings.Split(val, arraySep)
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
}

func toStringSlice(val string) ([]string, error) {
	split := strings.Split(val, arraySep)
	typedArray := make([]string, len(split))
	for i, raw := range split {
		typedArray[i] = raw
	}
	return typedArray, nil
}

func cast(typ, raw string) (interface{}, error) {
	switch typ {
	case "bigint":
		return toInt64(raw)
	case "bigint:array":
		return toInt64Slice(raw)
	case "bigint:pq-array":
		val, err := toInt64Slice(raw)
		if err != nil {
			return nil, err
		}
		return pq.Array(val), err
	case "int":
		return toInt32(raw)
	case "int:array":
		return toInt32Slice(raw)
	case "int:pq-array":
		val, err := toInt32Slice(raw)
		if err != nil {
			return nil, err
		}
		return pq.Array(val), err
	case "float":
		return toFloat64(raw)
	case "float:array":
		return toFloat64Slice(raw)
	case "float:pq-array":
		val, err := toFloat64Slice(raw)
		if err != nil {
			return nil, err
		}
		return pq.Array(val), err
	case "text":
		return raw, nil
	case "text:array":
		return toStringSlice(raw)
	case "text:pq-array":
		val, err := toStringSlice(raw)
		if err != nil {
			return nil, err
		}
		return pq.Array(val), err
	case "":
		return raw, nil
	default:
		return nil, fmt.Errorf("unknown data type [%s]", typ)
	}
}
