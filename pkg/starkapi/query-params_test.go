package starkapi

import (
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryParams_HashKey(t *testing.T) {
	p := QueryParams{}
	assert.Equal(t, "--------------0-0---", p.HashKey())

	assert.Equal(t, "--------------0-0---", p.HashKey())
	p.RequestName = "R"
	assert.Equal(t, "R--------------0-0---", p.HashKey())
	p.Id = "1"
	assert.Equal(t, "R-1-------------0-0---", p.HashKey())
	p.Ref = "2"
	assert.Equal(t, "R-1-2------------0-0---", p.HashKey())
	p.SiteId = "3"
	assert.Equal(t, "R-1-2-3-----------0-0---", p.HashKey())
	p.SiteRef = "4"
	assert.Equal(t, "R-1-2-3-4----------0-0---", p.HashKey())
	p.EquipRef = "5"
	assert.Equal(t, "R-1-2-3-4-5---------0-0---", p.HashKey())
	p.RuleName = "6"
	assert.Equal(t, "R-1-2-3-4-5-6--------0-0---", p.HashKey())
	p.RuleId = "7"
	assert.Equal(t, "R-1-2-3-4-5-6-7-------0-0---", p.HashKey())
	p.Severity = "8"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8------0-0---", p.HashKey())
	p.Duration = "9"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-----0-0---", p.HashKey())
	p.PersonId = "10"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10----0-0---", p.HashKey())
	p.Ts = "11"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11---0-0---", p.HashKey())
	p.EndTs = "12"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12--0-0---", p.HashKey())
	p.EventRef = "13"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12-13-0-0---", p.HashKey())
	p.Limit = 14
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12-13-14-0---", p.HashKey())
	p.Offset = 15
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12-13-14-15---", p.HashKey())
	p.SortA = "ts"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12-13-14-15-ts--", p.HashKey())
	p.SortD = "id"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12-13-14-15-ts-id-", p.HashKey())
	p.ProfileRef = "16"
	assert.Equal(t, "R-1-2-3-4-5-6-7-8-9-10-11-12-13-14-15-ts-id-16", p.HashKey())
}
func TestQueryParams_DecodeParameters_WithDefaultOperator(t *testing.T) {
	p := QueryParams{Id: "1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithEqual(t *testing.T) {
	p := QueryParams{Id: "<eq>1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithNotEqual(t *testing.T) {
	p := QueryParams{Id: "<nq>1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithGreaterThan(t *testing.T) {
	p := QueryParams{Id: "<gt>1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, ">", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithGreaterThanEqual(t *testing.T) {
	p := QueryParams{Id: "<ge>1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, ">=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithLessThan(t *testing.T) {
	p := QueryParams{Id: "<lt>1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "<", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeParameters_WithLessThanEqual(t *testing.T) {
	p := QueryParams{Id: "<le>1"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "<=", parameters[0].Operator)
	assert.Equal(t, int64(1), parameters[0].Value)
}

func TestQueryParams_DecodeAll_Case1(t *testing.T) {
	p := QueryParams{SiteRef: "<eq>s.abc", EquipId: "<gt>100", ProfileRef: "<eq>p.123"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(parameters))

	assert.Equal(t, "site_ref", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, "s.abc", parameters[0].Value)

	assert.Equal(t, "profile_ref", parameters[1].Column)
	assert.Equal(t, "=", parameters[1].Operator)
	assert.Equal(t, "p.123", parameters[1].Value)

	assert.Equal(t, "equip_id", parameters[2].Column)
	assert.Equal(t, ">", parameters[2].Operator)
	assert.Equal(t, int64(100), parameters[2].Value)
}

func TestQueryParams_DecodeAll_Case2(t *testing.T) {
	p := QueryParams{SiteRef: "<eq>s.abc", Id: "<nq>100"}
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

func TestQueryParams_build_sql(t *testing.T) {
	p := QueryParams{SiteRef: "<eq>s.abc", Id: "<nq>100", Ts: "1666797079", EndTs: "1666797080", ProfileRef: "<eq>p.123"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(100), parameters[0].Value)

	assert.Equal(t, "site_ref", parameters[1].Column)
	assert.Equal(t, "=", parameters[1].Operator)
	assert.Equal(t, "s.abc", parameters[1].Value)

	assert.Equal(t, "profile_ref", parameters[2].Column)
	assert.Equal(t, "=", parameters[2].Operator)
	assert.Equal(t, "p.123", parameters[2].Value)

	assert.Equal(t, "ts", parameters[3].Column)
	assert.Equal(t, "=", parameters[3].Operator)
	assert.Equal(t, int64(1666797079), parameters[3].Value)

	assert.Equal(t, "end_ts", parameters[4].Column)
	assert.Equal(t, "=", parameters[4].Operator)
	assert.Equal(t, int64(1666797080), parameters[4].Value)

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")

	assert.Equal(t, "Select * from hello where id != $1 and site_ref = $2 and profile_ref = $3 and ts = to_timestamp($4) and end_ts = to_timestamp($5)", sql)
	assert.Equal(t, 5, len(args))

}

func TestQueryParams_build_sql_SortA(t *testing.T) {
	p := QueryParams{SiteRef: "<eq>s.abc", Id: "<nq>100", Ts: "1666797079", EndTs: "1666797080", SortA: "endTs"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(100), parameters[0].Value)

	assert.Equal(t, "site_ref", parameters[1].Column)
	assert.Equal(t, "=", parameters[1].Operator)
	assert.Equal(t, "s.abc", parameters[1].Value)

	assert.Equal(t, "ts", parameters[2].Column)
	assert.Equal(t, "=", parameters[2].Operator)
	assert.Equal(t, int64(1666797079), parameters[2].Value)

	assert.Equal(t, "end_ts", parameters[3].Column)
	assert.Equal(t, "=", parameters[3].Operator)
	assert.Equal(t, int64(1666797080), parameters[3].Value)

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")

	assert.Equal(t, "Select * from hello where id != $1 and site_ref = $2 and ts = to_timestamp($3) and end_ts = to_timestamp($4) order by end_ts asc", sql)
	assert.Equal(t, 4, len(args))
}

func TestQueryParams_build_sql_SortD(t *testing.T) {
	p := QueryParams{SiteRef: "<eq>s.abc", Id: "<nq>100", Ts: "1666797079", EndTs: "1666797080", SortD: "endTs"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(100), parameters[0].Value)

	assert.Equal(t, "site_ref", parameters[1].Column)
	assert.Equal(t, "=", parameters[1].Operator)
	assert.Equal(t, "s.abc", parameters[1].Value)

	assert.Equal(t, "ts", parameters[2].Column)
	assert.Equal(t, "=", parameters[2].Operator)
	assert.Equal(t, int64(1666797079), parameters[2].Value)

	assert.Equal(t, "end_ts", parameters[3].Column)
	assert.Equal(t, "=", parameters[3].Operator)
	assert.Equal(t, int64(1666797080), parameters[3].Value)

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")

	assert.Equal(t, "Select * from hello where id != $1 and site_ref = $2 and ts = to_timestamp($3) and end_ts = to_timestamp($4) order by end_ts desc", sql)
	assert.Equal(t, 4, len(args))
}

func TestQueryParams_build_sql_SortAandSortD(t *testing.T) {
	p := QueryParams{SiteRef: "<eq>s.abc", Id: "<nq>100", Ts: "1666797079", EndTs: "1666797080", SortA: "endTs", SortD: "endTs"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(parameters))

	assert.Equal(t, "id", parameters[0].Column)
	assert.Equal(t, "!=", parameters[0].Operator)
	assert.Equal(t, int64(100), parameters[0].Value)

	assert.Equal(t, "site_ref", parameters[1].Column)
	assert.Equal(t, "=", parameters[1].Operator)
	assert.Equal(t, "s.abc", parameters[1].Value)

	assert.Equal(t, "ts", parameters[2].Column)
	assert.Equal(t, "=", parameters[2].Operator)
	assert.Equal(t, int64(1666797079), parameters[2].Value)

	assert.Equal(t, "end_ts", parameters[3].Column)
	assert.Equal(t, "=", parameters[3].Operator)
	assert.Equal(t, int64(1666797080), parameters[3].Value)

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")

	assert.Equal(t, "Select * from hello where id != $1 and site_ref = $2 and ts = to_timestamp($3) and end_ts = to_timestamp($4) order by end_ts asc", sql)
	assert.Equal(t, 4, len(args))
}

func TestQueryParams_DecodeProfileRef(t *testing.T) {
	p := QueryParams{ProfileRef: "<eq>p.123"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "profile_ref", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, "p.123", parameters[0].Value)
}

func TestQueryParams_WithIn(t *testing.T) {
	p := QueryParams{Severity: "<in>1,2,3"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parameters))

	assert.Equal(t, "severity", parameters[0].Column)
	assert.Equal(t, "IN", parameters[0].Operator)

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")

	assert.Equal(t, "Select * from hello where severity = ANY($1)", sql)
	assert.Equal(t, 1, len(args))

}

func TestQueryParams_WithInAndEqual(t *testing.T) {
	p := QueryParams{Severity: "<in>1,2,3", ProfileRef: "p.123"}
	parameters, err := p.DecodeParameters()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(parameters))

	assert.Equal(t, "profile_ref", parameters[0].Column)
	assert.Equal(t, "=", parameters[0].Operator)
	assert.Equal(t, "p.123", parameters[0].Value)

	assert.Equal(t, "severity", parameters[1].Column)
	assert.Equal(t, "IN", parameters[1].Operator)

	v := parameters[1].Value.(*pq.Int32Array)
	assert.NotNil(t, v)

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")

	assert.Equal(t, "Select * from hello where profile_ref = $1 and severity = ANY($2)", sql)
	assert.Equal(t, 2, len(args))
	assert.Equal(t, "p.123", args[0])

}

func TestQueryParams_WithInAndEqualAndSortAndLimitAndOffset(t *testing.T) {
	p := QueryParams{Severity: "<in>1,2,3", ProfileRef: "p.123", RuleName: "<in>ruleA,ruleB", SortD: "ts", Limit: 100, Offset: 10}

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")
	assert.Nil(t, err)

	assert.Equal(t, "Select * from hello where profile_ref = $1 and rule_name = ANY($2) and severity = ANY($3) order by ts desc LIMIT 100 OFFSET 10", sql)
	assert.Equal(t, 3, len(args))
	assert.Equal(t, "p.123", args[0])
}

func TestQueryParams_DateCreated(t *testing.T) {
	p := QueryParams{DateCreated: "1683731908"}

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")
	assert.Nil(t, err)

	assert.Equal(t, "Select * from hello where date_created = $1", sql)
	assert.Equal(t, 1, len(args))
	assert.Equal(t, int64(1683731908), args[0])
}

func TestQueryParams_OrderByDateCreated(t *testing.T) {
	p := QueryParams{Severity: "1", SortA: "dateCreated"}

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")
	assert.Nil(t, err)

	assert.Equal(t, "Select * from hello where severity = $1 order by date_created asc", sql)
	assert.Equal(t, 1, len(args))
	assert.Equal(t, int32(1), args[0])
}

func TestQueryParams_IssueStatus(t *testing.T) {
	p := QueryParams{IssueStatus: "1"}

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")
	assert.Nil(t, err)

	assert.Equal(t, "Select * from hello where issue_status_id = $1", sql)
	assert.Equal(t, 1, len(args))
	assert.Equal(t, int64(1), args[0])
}

func TestQueryParams_OrderByIssueStatus(t *testing.T) {
	p := QueryParams{SortA: "issueStatus"}

	sql, args, err := p.BuildParameterizedQuery("Select * from hello")
	assert.Nil(t, err)

	assert.Equal(t, "Select * from hello order by issue_status_id asc", sql)
	assert.Equal(t, 0, len(args))
}
