package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseContextPersonId_WithString(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, personId, "100")

	expectedId := int64(100)
	actualId, err := ParseContextPersonId(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedId, actualId)

}

func TestParseContextPersonId_WithFloat64(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, personId, float64(100))

	expectedId := int64(100)
	actualId, err := ParseContextPersonId(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedId, actualId)

}

func TestParseContextPersonId_WithInt64(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, personId, int64(100))

	expectedId := int64(100)
	actualId, err := ParseContextPersonId(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedId, actualId)

}

func TestParseContextPersonId_WithInt(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, personId, int(100))

	expectedId := int64(100)
	actualId, err := ParseContextPersonId(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedId, actualId)

}

func TestParseContextPersonId_WithInt32(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, personId, int32(100))

	expectedId := int64(100)
	actualId, err := ParseContextPersonId(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expectedId, actualId)

}

func TestParseContextPersonId_WithNone(t *testing.T) {

	ctx := context.Background()
	_, err := ParseContextPersonId(ctx)
	assert.NotNil(t, err)

}
