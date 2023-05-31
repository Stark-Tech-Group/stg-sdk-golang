package context

import (
	"context"
	"errors"
	"strconv"
)

const (
	personId = "personId"
)

func ParseContextPersonId(ctx context.Context) (int64, error) {

	val := ctx.Value(personId)
	var id int64
	switch val.(type) {
	case bool:
		return id, errors.New("unsupported type conversion bool -> int64")
	case int:
		return int64(val.(int)), nil
	case int32:
		return int64(val.(int32)), nil
	case int64:
		return val.(int64), nil
	case float32:
		return int64(val.(float32)), nil
	case float64:
		return int64(val.(float64)), nil
	case string:
		return strconv.ParseInt(val.(string), 10, 64)
	}

	return id, errors.New("unsupported type conversion")
}
