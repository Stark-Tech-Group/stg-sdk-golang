package starkapi

import (
	"fmt"
	"strings"
)

type arrayWithNull struct {
	values  []interface{}
	hasNull bool
	isNull  bool
}

func (arr arrayWithNull) toSql(seedIndex int, column string) string {
	if arr.hasNull {
		if arr.isNull {
			if len(arr.values) > 0 {
				return fmt.Sprintf("(%s = ANY($%d) or %s IS "+nullSql+")", column, seedIndex, column)
			} else {
				return fmt.Sprintf("%s IS "+nullSql, column)
			}
		} //TODO: add !arr.isNull functionality
	}

	return fmt.Sprintf("%s = ANY($%d)", column, seedIndex)
}

func toInt64Arr(val string) ([]interface{}, error) {
	split := strings.Split(val, arraySep)
	typedArray := make([]interface{}, len(split))
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

func toInt32Arr(val string) ([]interface{}, error) {
	split := strings.Split(val, arraySep)
	typedArray := make([]interface{}, len(split))
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

func toFloat64Arr(val string) ([]interface{}, error) {
	split := strings.Split(val, arraySep)
	typedArray := make([]interface{}, len(split))
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

func toStringArr(val string) ([]interface{}, error) {
	split := strings.Split(val, arraySep)
	typedArray := make([]interface{}, len(split))
	for i, raw := range split {
		typedArray[i] = raw
	}
	return typedArray, nil
}

func newArrayWithNull(raw string, typ string) (arrayWithNull, error) {
	arr := arrayWithNull{
		isNull:  false,
		hasNull: false,
		values:  make([]interface{}, 0),
	}
	var err error
	if strings.Contains(raw, nullVal) {
		split := strings.Split(raw, arraySep)
		for i := 0; i < len(split); i++ {
			if split[i] == nullVal {
				split = append(split[:i], split[i+1:]...)
				break
			}
		}
		arr.hasNull = true
		arr.isNull = true // will need to handle not null in future, isNull = false == NOT NULL
		raw = strings.Join(split, arraySep)
		if raw == "" {
			return arr, nil
		}
	}
	switch typ {
	case "bigint:array":
		arr.values, err = toInt64Arr(raw)
	case "bigint:pq-array":
		arr.values, err = toInt64Arr(raw)
	case "int:array":
		arr.values, err = toInt32Arr(raw)
	case "int:pq-array":
		arr.values, err = toInt32Arr(raw)
	case "float:array":
		arr.values, err = toFloat64Arr(raw)
	case "float:pq-array":
		arr.values, err = toFloat64Arr(raw)
	case "text:array":
		arr.values, err = toStringArr(raw)
	case "text:pq-array":
		arr.values, err = toStringArr(raw)
	default:
		return arr, fmt.Errorf("unknown data type [%s]", typ)
	}

	return arr, err
}
