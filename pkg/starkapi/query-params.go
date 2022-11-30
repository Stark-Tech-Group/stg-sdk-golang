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
	in              = "IN"
)

var operatorMap = map[string]string{
	"<eq>": "=",
	"<nq>": "!=",
	"<gt>": ">",
	"<ge>": ">=",
	"<lt>": "<",
	"<le>": "<=",
	"<in>": in,
}
var input = regexp.MustCompile("^([a-z]|[A-Z]|[0-9]|[.]|-){1,75}$")
var colTypeMap map[string]*reflect.StructField
var fieldToColumnMap map[string]string

type QueryParams struct {
	Id          string `json:"id" schema:"id" sqlColumn:"id" sqlType:"bigint"`
	Ref         string `json:"ref" schema:"ref" sqlColumn:"ref" sqlType:"text"`
	SiteId      string `json:"siteId" schema:"siteId" sqlColumn:"site_id" sqlType:"bigint"`
	SiteRef     string `json:"siteRef" schema:"siteRef" sqlColumn:"site_ref" sqlType:"text"`
	ProfileRef  string `json:"profileRef" schema:"profileRef" sqlColumn:"profile_ref" sqlType:"text"`
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
	return fmt.Sprintf("%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%s-%v-%v-%s-%s-%s",
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
		q.SortD,
		q.ProfileRef)
}

func (q *QueryParams) Validate() bool {
	if q.Limit > maxLimit {
		return false
	}

	return true
}

func castWithColumn(column, raw string) (interface{}, error) {
	field := getSqlTypeByColumnName(column)
	if field == nil {
		return "", fmt.Errorf("field not found with sql column name [%s]", column)
	}
	return castWithField(field, raw)
}

func castWithField(field *reflect.StructField, raw string) (interface{}, error) {
	sqlValType := field.Tag.Get(sqlType)
	switch sqlValType {
	case "bigint":
		value, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			logger.Errorf("faild to convert strign to int64")
			return nil, errors.New("not a int64")
		}
		return value, nil
	case "int":
		value, err := strconv.ParseInt(raw, 10, 32)
		if err != nil {
			logger.Errorf("faild to convert string to int32")
			return nil, errors.New("not a int32")
		}
		return value, nil
	case "float":
		value, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			logger.Errorf("faild to convert string to float64")
			return nil, errors.New("not a float64")
		}
		return value, nil
	case "text":
		return raw, nil
	case "":
		return raw, nil
	default:
		return nil, fmt.Errorf("unknown data type [%s]", sqlValType)
	}
}

func decodeRightSide(field *reflect.StructField, val string) (string, interface{}, error) {

	var operator, raw string

	if val[0:1] == leftOp && len(val) > 4 {
		queryOp := val[0:4]
		operator = operatorMap[queryOp]
		raw = val[4:]

		if operator == in {
			return operator, raw, nil
		}

	} else {
		operator = defaultOperator
		raw = val
	}

	if len(operator) == 0 {
		logger.Errorf("no operator found while decoding query to sql")
		return "", "", errors.New("no operator found")
	}

	value, err := castWithField(field, raw)
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
			if field.Name == "SortA" || field.Name == "SortD" {
				sqlTag := getColumnNameByFieldName(val)
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

//scan preforms a reflection lookup to populate internal collections for faster lookups
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

func getSqlTypeByColumnName(column string) *reflect.StructField {
	scan()
	return colTypeMap[column]
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
		b.WriteString(where)
	}

	explodedIndex := 0
	for i, p := range parameters {
		if p.AscSort == false && p.DescSort == false {
			chunk, explodedArgs := p.parameterizedClause(i + explodedIndex)

			b.WriteString(chunk)
			if explodedArgs != nil {
				for _, v := range explodedArgs {
					args = append(args, v)
				}
				explodedIndex += len(explodedArgs)
			} else {
				args = append(args, p.Value)
			}

			//evaluates the position current index for 'order by' and 'and'
			if i < len(parameters)-1 && !parameters[i+1].AscSort && !parameters[i+1].DescSort {
				b.WriteString(and)
			} else if i < len(parameters)-1 && (parameters[i+1].AscSort || parameters[i+1].DescSort) {
				b.WriteString(orderBy)
			}

		} else if p.AscSort || p.DescSort {
			chunk, _ := p.parameterizedClause(i + explodedIndex)
			b.WriteString(chunk)
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

func (p *Parameter) parameterizedClause(seedIndex int) (string, []interface{}) {

	if p.Operator == in {
		return p.parameterizedInClause(seedIndex + 1)

	} else {
		val := fmt.Sprintf("$%d", seedIndex+1)
		if p.Decorator != "" {
			val = strings.Replace(p.Decorator, "%", fmt.Sprintf("$%d", seedIndex+1), 1)
		}
		if p.AscSort {
			return fmt.Sprintf("%s asc", p.Value), nil
		} else if p.DescSort {
			return fmt.Sprintf("%s desc", p.Value), nil
		}
		return fmt.Sprintf("%s %s %s", p.Column, p.Operator, val), nil
	}
}

func (p *Parameter) parameterizedInClause(num int) (string, []interface{}) {
	values := strings.Split(p.Value.(string), ",")
	builder := strings.Builder{}

	explodedArgs := make([]interface{}, len(values))
	builder.WriteString("(")
	for i, v := range values {
		builder.WriteString(fmt.Sprintf("$%d", num+i))
		value, err := castWithColumn(p.Column, v)
		if err == nil {
			explodedArgs[i] = value
		}
		if i < len(values)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString(")")

	return fmt.Sprintf("%s %s %s", p.Column, p.Operator, builder.String()), explodedArgs
}
