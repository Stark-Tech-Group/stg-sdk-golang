package starkapi

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
)

const (
	maxLimit        = 5000
	sqlColumn       = "sqlColumn"
	sqlType         = "sqlType"
	sqlDecorator    = "sqlDecorator"
	where           = " where "
	and             = " and "
	defaultOperator = "="
	leftOp          = "<"
	schema          = "schema"
	orderBy         = " order by "
	in              = "IN"
	pqArrayType     = ":pq-array"
	startLike       = "like %"
	endLike         = "% like"
	nullSql         = "NULL"
	nullVal         = "null"
)

var operatorMap = map[string]string{
	"<eq>": "=",
	"<nq>": "!=",
	"<gt>": ">",
	"<ge>": ">=",
	"<lt>": "<",
	"<le>": "<=",
	"<in>": in,
	"<sw>": startLike,
	"<ew>": endLike,
}
var input = regexp.MustCompile("^([a-z]|[A-Z]|[0-9]|[.]|-){1,75}$")
var colTypeMap map[string]*reflect.StructField
var fieldToColumnMap map[string]string

type QueryParams struct {
	Id            string `json:"id" schema:"id" sqlColumn:"id" sqlType:"bigint"`
	Ref           string `json:"ref" schema:"ref" sqlColumn:"ref" sqlType:"text"`
	ProfileName   string `json:"profileName" schema:"profileName" sqlColumn:"profile_name" sqlType:"text"`
	SiteId        string `json:"siteId" schema:"siteId" sqlColumn:"site_id" sqlType:"bigint"`
	SiteRef       string `json:"siteRef" schema:"siteRef" sqlColumn:"site_ref" sqlType:"text"`
	SiteName      string `json:"siteName" schema:"siteName" sqlColumn:"site_name" sqlType:"text"`
	ProfileRef    string `json:"profileRef" schema:"profileRef" sqlColumn:"profile_ref" sqlType:"text"`
	EquipId       string `json:"equipId" schema:"equipId" sqlColumn:"equip_id" sqlType:"bigint"`
	EquipRef      string `json:"equipRef" schema:"equipRef" sqlColumn:"equip_ref" sqlType:"text"`
	EquipName     string `json:"equipName" schema:"equipName" sqlColumn:"equip_name" sqlType:"text"`
	RuleName      string `json:"ruleName" schema:"ruleName" sqlColumn:"rule_name" sqlType:"text"`
	RuleId        string `json:"ruleId" schema:"ruleId" sqlColumn:"rule_id" sqlType:"bigint"`
	Severity      string `json:"severity" schema:"severity" sqlColumn:"severity" sqlType:"int"`
	Duration      string `json:"dur" schema:"dur" sqlColumn:"dur" sqlType:"bigint"`
	PersonId      string `json:"personId" schema:"personId" sqlColumn:"person_id" sqlType:"bigint"`
	Ts            string `json:"ts" schema:"ts" sqlColumn:"ts" sqlType:"bigint" sqlDecorator:"to_timestamp(%)"`
	EndTs         string `json:"endTs" schema:"endTs" sqlColumn:"end_ts" sqlType:"bigint" sqlDecorator:"to_timestamp(%)"`
	EventRef      string `json:"eventRef" schema:"eventRef" sqlColumn:"event_ref" sqlType:"text"`
	Name          string `json:"name" schema:"name" sqlColumn:"name" sqlType:"text"`
	GeoCity       string `json:"geoCity" schema:"geoCity" sqlColumn:"geo_city" sqlType:"text"`
	GeoPostalCode string `json:"geoPostalCode" schema:"geoPostalCode" sqlColumn:"geo_postal_code" sqlType:"text"`
	Lat           string `json:"lat" schema:"lat" sqlColumn:"lat" sqlType:"float"`
	Lon           string `json:"lon" schema:"lon" sqlColumn:"lon" sqlType:"float"`
	Gsf           string `json:"gsf" schema:"gsf" sqlColumn:"gsf" sqlType:"int"`
	GeoAddress2   string `json:"geoAddress2" schema:"geoAddress2" sqlColumn:"geo_address2" sqlType:"text"`
	GeoAddress1   string `json:"geoAddress1" schema:"geoAddress1" sqlColumn:"geo_address1" sqlType:"text"`
	GeoStateCode  string `json:"geoStateCode" schema:"geoStateCode" sqlColumn:"geo_state_code" sqlType:"text"`
	Description   string `json:"description" schema:"description" sqlColumn:"description" sqlType:"text"`

	Limit         int    `json:"limit" schema:"limit"`
	Offset        int    `json:"offset" schema:"offset"`
	RequestName   string `json:"-" schema:"-"`
	EventType     string `json:"eventType" schema:"eventType" sqlColumn:"event_type" sqlType:"text"`
	DateCreated   string `json:"dateCreated" schema:"dateCreated" sqlColumn:"date_created" sqlType:"bigint"`
	IssueStatus   string `json:"issueStatus" schema:"issueStatus" sqlColumn:"issue_status_id" sqlType:"bigint"`
	TargetRef     string `json:"targetRef" schema:"targetRef" sqlColumn:"target_ref" sqlType:"text"`
	EquipType     string `json:"equipType" schema:"equipType" sqlColumn:"equip_type_name" sqlType:"text"`
	EquipTypeName string `json:"equipTypeName" schema:"equipTypeName" sqlColumn:"equip_type_name" sqlType:"text"`
	EquipTypeId   string `json:"equipTypeId" schema:"equipTypeId" sqlColumn:"equip_type_id" sqlType:"bigint"`
	PointTypeName string `json:"pointTypeName" schema:"pointTypeName" sqlColumn:"point_type_name" sqlType:"text"`
	PointTypeId   string `json:"pointTypeId" schema:"pointTypeId" sqlColumn:"point_type_id" sqlType:"bigint"`
	PointUridName string `json:"pointUridName" schema:"pointUridName" sqlColumn:"point_urid_name" sqlType:"text"`
	PointUridId   string `json:"pointUridId" schema:"pointUridId" sqlColumn:"point_urid_id" sqlType:"bigint"`
	Unit          string `json:"unit" schema:"unit" sqlColumn:"unit" sqlType:"text"`
	Category      string `json:"category" schema:"category" sqlColumn:"category" sqlType:"text"`
	Batch         string `json:"batch" schema:"batch" sqlColumn:"batch" sqlType:"text"`
	SortA         string `json:"sortA" schema:"sortA"`
	SortD         string `json:"sortD" schema:"sortD"`
}

// HashKey creates a compounded string of the current QueryParams
func (q *QueryParams) HashKey() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%v-%v-%s-%s-%s",
		q.RequestName,
		q.Id,
		q.Ref,
		q.SiteId,
		q.SiteRef,
		q.EquipRef,
		q.RuleName,
		q.RuleId,
		q.Severity,
		q.Duration,
		q.PersonId,
		q.Ts,
		q.EndTs,
		q.EventRef,
		q.Limit,
		q.Offset,
		q.SortA,
		q.SortD,
		q.ProfileRef)
}

func (q *QueryParams) Validate() bool {
	if q.Limit > maxLimit {
		return false
	}

	return true
}

func castWithField(field *reflect.StructField, raw string, operator string) (interface{}, error) {
	sqlValType := field.Tag.Get(sqlType)

	if operator == in {
		sqlValType = sqlValType + pqArrayType
		arr, err := newArrayWithNull(raw, sqlValType)
		if err != nil {
			logger.Errorf("error converting to array with null: %s", err)
			return nil, err
		}
		return arr, nil
	}

	return cast(sqlValType, raw)
}

func decodeRightSide(field *reflect.StructField, val string) (string, interface{}, error) {

	var operator, raw string

	var value interface{}

	if val[0:1] == leftOp && len(val) > 4 {
		queryOp := val[0:4]
		operator = operatorMap[queryOp]
		raw = val[4:]
	} else {
		operator = defaultOperator
		raw = val
	}

	if len(operator) == 0 {
		logger.Errorf("no operator found while decoding query to sql")
		return "", "", errors.New("no operator found")
	}

	if strings.Contains(val, nullVal) && operator != in {
		value = nullVal
		return operator, value, nil
	}

	value, err := castWithField(field, raw, operator)
	if err != nil {
		return "", nil, err
	}

	return operator, value, nil
}

func (q *QueryParams) DecodeParameters() ([]Parameter, error) {
	scan()
	t := reflect.TypeOf(q).Elem()
	value := reflect.Indirect(reflect.ValueOf(q))

	sorted := false

	clauses := make([]Parameter, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := value.FieldByName(field.Name).String()
		decorator := field.Tag.Get(sqlDecorator)
		if len(val) > 0 {
			typ := field.Tag.Get(sqlType)

			if field.Name == "SortA" || field.Name == "SortD" {
				sqlTag := getColumnNameByFieldName(val)
				operator, sqlValue, err := decodeRightSide(&field, sqlTag)
				if err != nil {
					return nil, err
				}
				if field.Name == "SortA" && !sorted {
					clauses = append(clauses, Parameter{Column: sqlTag, Operator: operator, Value: sqlValue,
						Decorator: decorator, AscSort: true, DescSort: false, Type: typ})
					sorted = true
				} else if field.Name == "SortD" && !sorted {
					clauses = append(clauses, Parameter{Column: sqlTag, Operator: operator, Value: sqlValue,
						Decorator: decorator, AscSort: false, DescSort: true, Type: typ})
					sorted = true
				}
			}
			tag := field.Tag.Get(sqlColumn)
			if len(tag) > 0 {
				operator, sqlValue, err := decodeRightSide(&field, val)
				if sqlValue == nullVal {
					typ = "text"
				}
				if err != nil {
					return nil, err
				}
				clauses = append(clauses, Parameter{Column: tag, Operator: operator, Value: sqlValue,
					Decorator: decorator, AscSort: false, DescSort: false, Type: typ})
			}
		}
	}

	return clauses, nil
}

func (q *QueryParams) findSqlColumn(t reflect.Type, sortVal string) string {
	sqlTag := ""
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get(schema) == sortVal {
			sqlTag = field.Tag.Get(sqlColumn)
		}
	}
	return sqlTag
}

// scan preforms a reflection lookup to populate internal collections for faster lookups
func scan() {
	if colTypeMap == nil || fieldToColumnMap == nil {

		t := reflect.ValueOf(&QueryParams{}).Elem()

		colTypeMap = make(map[string]*reflect.StructField, 0)
		fieldToColumnMap = make(map[string]string, 0)
		for i := 0; i < t.NumField(); i++ {
			field := t.Type().Field(i)
			col := field.Tag.Get(sqlColumn)
			sch := field.Tag.Get(schema)
			if len(col) > 0 {
				colTypeMap[col] = &field
			}
			if len(sch) > 0 && len(field.Tag.Get(sqlColumn)) > 0 {
				fieldToColumnMap[sch] = field.Tag.Get(sqlColumn)
			}
		}
	}
}

func getColumnNameByFieldName(field string) string {
	scan()
	return fieldToColumnMap[field]

}

// BuildParameterizedQuery appends a parametrized query to the provided sql statement and returns the query with arguments
func (q *QueryParams) BuildParameterizedQuery(sql string) (string, []interface{}, error) {
	parameters, err := q.DecodeParameters()
	if err != nil {
		return "", nil, errors.New("bad parameters")
	}

	args := make([]interface{}, 0)

	b := strings.Builder{}
	b.WriteString(sql)

	if len(parameters) > 0 {
		if len(parameters) == 1 && (parameters[0].AscSort || parameters[0].DescSort) {
			b.WriteString(orderBy)
		} else if parameters[0].AscSort || parameters[0].DescSort {
			val := parameters[len(parameters)-1]
			parameters[len(parameters)-1] = parameters[0]
			parameters[0] = val
			b.WriteString(where)
		} else {
			b.WriteString(where)
		}
	}

	explodedIndex := 0
	for i, p := range parameters {
		if p.AscSort == false && p.DescSort == false {
			chunk := p.parameterizedClause(i + explodedIndex)

			b.WriteString(chunk)
			if p.Operator == in {
				if len(p.Value.(arrayWithNull).values) > 0 {
					args = append(args, pq.Array(p.Value.(arrayWithNull).values))
				}
			} else if p.Value != nil && p.Value != nullVal {
				args = append(args, p.Value)
			} else {
				explodedIndex--
			}

			//evaluates the position current index for 'order by' and 'and'
			if i < len(parameters)-1 && !parameters[i+1].AscSort && !parameters[i+1].DescSort {
				b.WriteString(and)
			} else if i < len(parameters)-1 && (parameters[i+1].AscSort || parameters[i+1].DescSort) {
				b.WriteString(orderBy)
			}

		} else if p.AscSort || p.DescSort {
			chunk := p.parameterizedClause(i + explodedIndex)
			b.WriteString(chunk)

			if i < len(parameters)-1 && !parameters[i+1].AscSort && !parameters[i+1].DescSort {
				b.WriteString(and)
			} else if i < len(parameters)-1 && (parameters[i+1].AscSort || parameters[i+1].DescSort) {
				b.WriteString(orderBy)
			}
		}
	}

	if q.Limit > 0 {
		b.WriteString(fmt.Sprintf(" LIMIT %v OFFSET %v", q.Limit, q.Offset))
	}

	return b.String(), args, nil
}

type Parameter struct {
	Column    string
	Operator  string
	Value     interface{}
	Type      string
	Decorator string
	AscSort   bool
	DescSort  bool
}

func (p *Parameter) parameterizedClause(seedIndex int) string {

	if p.Operator == in {
		return (p.Value.(arrayWithNull)).toSql(seedIndex+1, p.Column)
	} else if p.Operator == startLike {
		if p.nullValCheck() {
			return fmt.Sprintf("%s IS "+nullSql, p.Column)
		}
		p.Value = p.Value.(string) + "%"
		return fmt.Sprintf("%s like $%d", p.Column, seedIndex+1)
	} else if p.Operator == endLike {
		if p.nullValCheck() {
			return fmt.Sprintf("%s IS "+nullSql, p.Column)
		}
		p.Value = "%" + p.Value.(string)
		return fmt.Sprintf("%s like $%d", p.Column, seedIndex+1)
	} else {
		if p.nullValCheck() {
			return fmt.Sprintf("%s IS "+nullSql, p.Column)
		}
		val := fmt.Sprintf("$%d", seedIndex+1)
		if p.Decorator != "" {
			val = strings.Replace(p.Decorator, "%", fmt.Sprintf("$%d", seedIndex+1), 1)
		}
		if p.AscSort {
			return fmt.Sprintf("%s asc", p.Value)
		} else if p.DescSort {
			return fmt.Sprintf("%s desc", p.Value)
		}
		return fmt.Sprintf("%s %s %s", p.Column, p.Operator, val)
	}
}

func (p *Parameter) nullValCheck() bool {
	return p.Type == "text" && strings.Contains(p.Value.(string), nullVal)
}
