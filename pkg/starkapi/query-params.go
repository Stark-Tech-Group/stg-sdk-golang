package starkapi

import (
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strconv"
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
)

var operatorMap = map[string]string{
	"<eq>": "=",
	"<nq>": "!=",
	"<gt>": ">",
	"<ge>": ">=",
	"<lt>": "<",
	"<le>": "<=",
}
var input = regexp.MustCompile("^([a-z]|[A-Z]|[0-9]|[.]|-){1,75}$")

type QueryParams struct {
	Id          string `json:"id" schema:"id" sqlColumn:"id" sqlType:"bigint"`
	Ref         string `json:"ref" schema:"ref" sqlColumn:"ref" sqlType:"text"`
	SiteId      string `json:"siteId" schema:"siteId" sqlColumn:"site_id" sqlType:"bigint"`
	SiteRef     string `json:"siteRef" schema:"siteRef" sqlColumn:"site_ref" sqlType:"text"`
	EquipId     string `json:"equipId" schema:"equipId" sqlColumn:"equip_id" sqlType:"bigint"`
	EquipRef    string `json:"equipRef" schema:"equipRef" sqlColumn:"equip_ref" sqlType:"text"`
	RuleName    string `json:"ruleName" schema:"ruleName" sqlColumn:"rule_name" sqlType:"text"`
	RuleId      string `json:"ruleId" schema:"ruleId" sqlColumn:"rule_id" sqlType:"bigint"`
	Severity    string `json:"severity" schema:"severity" sqlColumn:"severity" sqlType:"int"`
	Duration    string `json:"dur" schema:"dur" sqlColumn:"dur" sqlType:"bigint"`
	PersonId    string `json:"personId" schema:"personId" sqlColumn:"person_id" sqlType:"bigint"`
	Ts          string `json:"ts" schema:"ts" sqlColumn:"ts" sqlType:"bigint" sqlDecorator:"to_timestamp(%)"`
	EndTs       string `json:"endTs" schema:"endTs" sqlColumn:"end_ts" sqlType:"bigint" sqlDecorator:"to_timestamp(%)"`
	Limit       int    `json:"limit" schema:"limit"`
	Offset      int    `json:"offset" schema:"offset"`
	RequestName string `json:"-" schema:"-"`
	EventType   string `json:"eventType" schema:"eventType" sqlColumn:"event_type" sqlType:"text"`
	SortA       string `json:"sortA" schema:"sortA"`
	SortD       string `json:"sortD" schema:"sortD"`
}

// HashKey creates a compounded string of the current QueryParams
func (q *QueryParams) HashKey() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%v-%v-%s-%s",
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
		q.Limit,
		q.Offset,
		q.SortA,
		q.SortD)
}

func (q *QueryParams) Validate() bool {
	if q.Limit > maxLimit {
		return false
	}

	return true
}

func decodeRightSide(field *reflect.StructField, val string) (string, interface{}, error) {

	var operator, raw string

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

	sqlValType := field.Tag.Get(sqlType)
	switch sqlValType {
	case "bigint":
		value, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			logger.Errorf("faild to convert strign to int64")
			return "", nil, errors.New("not a int64")
		}
		return operator, value, nil
	case "int":
		value, err := strconv.ParseInt(raw, 10, 32)
		if err != nil {
			logger.Errorf("faild to convert string to int32")
			return "", nil, errors.New("not a int32")
		}
		return operator, value, nil
	case "float":
		value, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			logger.Errorf("faild to convert string to float64")
			return "", nil, errors.New("not a float64")
		}
		return operator, value, nil
	default:
		return operator, raw, nil
	}

}

func (q *QueryParams) DecodeParameters() ([]Parameter, error) {

	t := reflect.TypeOf(q).Elem()
	value := reflect.Indirect(reflect.ValueOf(q))

	sorted := false

	clauses := make([]Parameter, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := value.FieldByName(field.Name).String()
		decorator := field.Tag.Get(sqlDecorator)
		if len(val) > 0 {
			if field.Name == "SortA" || field.Name == "SortD" {
				sqlTag := q.findSqlColumn(t, val)
				operator, sqlValue, err := decodeRightSide(&field, sqlTag)
				if err != nil {
					return nil, err
				}
				if field.Name == "SortA" && !sorted {
					clauses = append(clauses, Parameter{Column: sqlTag, Operator: operator, Value: sqlValue, Decorator: decorator, AscSort: true, DescSort: false})
					sorted = true
				} else if field.Name == "SortD" && !sorted {
					clauses = append(clauses, Parameter{Column: sqlTag, Operator: operator, Value: sqlValue, Decorator: decorator, AscSort: false, DescSort: true})
					sorted = true
				}
			}
			tag := field.Tag.Get(sqlColumn)
			if len(tag) > 0 {
				operator, sqlValue, err := decodeRightSide(&field, val)
				if err != nil {
					return nil, err
				}
				clauses = append(clauses, Parameter{Column: tag, Operator: operator, Value: sqlValue, Decorator: decorator, AscSort: false, DescSort: false})
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

// BuildParameterizedQuery appends a parametrized query to the provided sql statement and returns the query with arguments
func (q *QueryParams) BuildParameterizedQuery(sql string) (string, []interface{}, error) {
	parameters, err := q.DecodeParameters()
	if err != nil {
		return "", nil, errors.New("bad parameters")
	}

	args := make([]interface{}, len(parameters))

	b := strings.Builder{}
	b.WriteString(sql)
	if len(parameters) > 0 {
		b.WriteString(where)
	}
	for i, p := range parameters {
		if p.AscSort == false && p.DescSort == false {
			b.WriteString(p.parameterizedClause(i + 1))
			if i < len(parameters)-1 && !parameters[i+1].AscSort && !parameters[i+1].DescSort {
				b.WriteString(and)
			} else if i < len(parameters)-1 && (parameters[i+1].AscSort || parameters[i+1].DescSort) {
				b.WriteString(orderBy)
			}
			args[i] = p.Value
		} else if p.AscSort {
			b.WriteString(p.parameterizedClause(i + 1))
			args[i] = args[len(args)-1]
			args = args[:len(args)-1]
		} else if p.DescSort {
			b.WriteString(p.parameterizedClause(i + 1))
			args[i] = args[len(args)-1]
			args = args[:len(args)-1]
		}
	}

	if q.Limit > 0 {
		b.WriteString(fmt.Sprintf(" LIMIT %v OFFSET %v", q.Limit, q.Offset))
	} else {
		b.WriteString(fmt.Sprintf(" LIMIT %v", maxLimit))
	}

	return b.String(), args, nil
}

type Parameter struct {
	Column    string
	Operator  string
	Value     interface{}
	Decorator string
	AscSort   bool
	DescSort  bool
}

func (p *Parameter) parameterizedClause(num int) string {

	val := fmt.Sprintf("$%d", num)
	if p.Decorator != "" {
		val = strings.Replace(p.Decorator, "%", fmt.Sprintf("$%d", num), 1)
	}
	if p.AscSort {
		return fmt.Sprintf("%s asc", p.Value)
	} else if p.DescSort {
		return fmt.Sprintf("%s desc", p.Value)
	}
	return fmt.Sprintf("%s %s %s", p.Column, p.Operator, val)
}
