package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryParams_HashKey(t *testing.T) {
	p := QueryParams{}
	assert.Equal(t, "------------0-0", p.HashKey())
	p.Id = "1"
	assert.Equal(t, "1------------0-0", p.HashKey())
	p.Ref = "2"
	assert.Equal(t, "1-2-----------0-0", p.HashKey())
	p.SiteId = "3"
	assert.Equal(t, "1-2-3----------0-0", p.HashKey())
	p.SiteRef = "4"
	assert.Equal(t, "1-2-3-4---------0-0", p.HashKey())
	p.EquipRef = "5"
	assert.Equal(t, "1-2-3-4-5--------0-0", p.HashKey())
	p.RuleName = "6"
	assert.Equal(t, "1-2-3-4-5-6-------0-0", p.HashKey())
	p.RuleId = "7"
	assert.Equal(t, "1-2-3-4-5-6-7------0-0", p.HashKey())
	p.Severity = "8"
	assert.Equal(t, "1-2-3-4-5-6-7-8-----0-0", p.HashKey())
	p.Duration = "9"
	assert.Equal(t, "1-2-3-4-5-6-7-8-9----0-0", p.HashKey())
	p.PersonId = "10"
	assert.Equal(t, "1-2-3-4-5-6-7-8-9-10---0-0", p.HashKey())
	p.Ts = "11"
	assert.Equal(t, "1-2-3-4-5-6-7-8-9-10-11--0-0", p.HashKey())
	p.EndTs = "12"
	assert.Equal(t, "1-2-3-4-5-6-7-8-9-10-11-12-0-0", p.HashKey())
	p.Limit = 13
	assert.Equal(t, "1-2-3-4-5-6-7-8-9-10-11-12-13-0", p.HashKey())
	p.Offset = 14
	assert.Equal(t, "1-2-3-4-5-6-7-8-9-10-11-12-13-14", p.HashKey())
}

func TestQueryParams_DecodeParameters_WithEqual(t *testing.T) {
	p := QueryParams{Id: "[eq]1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithNotEqual(t *testing.T) {
	p := QueryParams{Id: "[nq]1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithGreaterThan(t *testing.T) {
	p := QueryParams{Id: "[gt]1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, ">", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithGreaterThanEqual(t *testing.T) {
	p := QueryParams{Id: "[ge]1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, ">=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithLessThan(t *testing.T) {
	p := QueryParams{Id: "[lt]1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "<", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithLessThanEqual(t *testing.T) {
	p := QueryParams{Id: "[le]1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "<=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeAll_Case1(t *testing.T) {
	p := QueryParams{SiteRef: "[eq]s.abc", EquipId: "[gt]100"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(parameters))

	assert.Equal(t, "site_ref", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, "s.abc", parameters[0].Value)

	assert.Equal(t, "equip_id", parameters[1].Column)
	assert.Equal(t, ">", parameters[1].Operator)
	assert.Equal(t, int64(100), parameters[1].Value)
}

func TestQueryParams_DecodeAll_Case2(t *testing.T) {
	p := QueryParams{SiteRef: "[eq]s.abc", Id: "[nq]100"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(100), parameters[0].Value)

	assert.Equal(t, "site_ref", parameters[1].Column)
	assert.Equal(t, "=", parameters[1].Operator)
	assert.Equal(t, "s.abc", parameters[1].Value)

}
