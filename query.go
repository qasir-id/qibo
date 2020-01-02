package qibo

import (
	"bytes"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Query struct binder for default query param
type Query struct {
	Page   int32  `json:"page"`
	Count  int32  `json:"count"`
	Sort   string `json:"sort"`
	Filter map[string]interface{}
}

// Operator string translation
var Operator = map[string]string{
	"gt":   ">",
	"lt":   "<",
	"eq":   "=",
	"ne":   "!=",
	"gte":  ">=",
	"lte":  "<=",
	"like": "like",
	"in":   "in",
}

func (q *Query) SetFilter(filter map[string]interface{}) {
	q.Filter = filter
}

func (q *Query) GetFilter() map[string]interface{} {
	return q.Filter
}

// Where generate sql WHERE statement ,with format
//		key :"{columnName}{$operator}"
//		value : interface
// with default operator value "$eq"
// for example :
//     "amount$gte": 19200.00
// 	   "status": 1
// will be translated into sql format :
// 		WHERE amount >= 19200.00
//		AND status = 1
func (q *Query) Where() (string, []interface{}) {
	query := new(bytes.Buffer)
	var args []interface{}
	i := 0
	for k, v := range q.Filter {
		var validDate = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
		fields := strings.Split(k, "$")
		columnName := fields[0]
		isRequire := func(s string) bool {
			return s[len(s)-1:] == "!"
		}(fields[1])
		opr := translateOperator(strings.TrimSuffix(fields[1], "!"))

		if i == 0 {
			if isRequire || !IsArgNil(v) {
				switch opr {
				case Operator["like"]:
					query.WriteString(` ` + columnName + ` ` + opr + ` ?`)
					tmpArgs, _ := v.(string)
					args = append(args, "%"+tmpArgs+"%")
				case Operator["in"]:
					query.WriteString(` ` + columnName + ` ` + opr + ` (?) `)
					args = append(args, v)
				case Operator["lte"]:
					tmpArgs, _ := v.(string)
					if validDate.MatchString(tmpArgs) {
						tmpArgs += " 23:59:59"
					}
					query.WriteString(` ` + columnName + ` ` + opr + ` ? `)
					args = append(args, tmpArgs)
				case Operator["gte"]:
					tmpArgs, _ := v.(string)
					if validDate.MatchString(tmpArgs) {
						tmpArgs += " 00:00:00"
					}
					query.WriteString(` ` + columnName + ` ` + opr + ` ? `)
					args = append(args, tmpArgs)
				default:
					query.WriteString(` ` + columnName + ` ` + opr + ` ? `)
					args = append(args, v)
				}
			} else {
				query.WriteString(` 1 = 1 `)
			}

		} else {
			if isRequire || !IsArgNil(v) {
				switch opr {
				case Operator["like"]:
					query.WriteString(` AND ` + columnName + ` ` + opr + ` ? `)
					tmpArgs, _ := v.(string)
					args = append(args, "%"+tmpArgs+"%")
				case Operator["in"]:
					query.WriteString(` AND ` + columnName + ` ` + opr + ` (?) `)
					args = append(args, v)
				case Operator["lte"]:
					tmpArgs, _ := v.(string)
					if validDate.MatchString(tmpArgs) {
						tmpArgs += " 23:59:59"
					}
					query.WriteString(` AND ` + columnName + ` ` + opr + ` ? `)
					args = append(args, tmpArgs)
				case Operator["gte"]:
					tmpArgs, _ := v.(string)
					if validDate.MatchString(tmpArgs) {
						tmpArgs += " 00:00:00"
					}
					query.WriteString(` AND ` + columnName + ` ` + opr + ` ? `)
					args = append(args, tmpArgs)
				default:
					query.WriteString(` AND ` + columnName + ` ` + opr + ` ? `)
					args = append(args, v)
				}
			} else {
				query.WriteString(` AND 1 = 1 `)
			}
		}
		i++
	}
	return query.String(), args
}

// Order generate string ordering query statement
func (q *Query) Order() string {
	if len(q.Sort) > 0 {
		field := strings.Split(q.Sort, ",")
		var sort string
		for _, v := range field {
			sortType := func(str string) string {
				if strings.HasPrefix(str, "-") {
					return `desc`
				}
				return `asc`
			}
			sort += strings.TrimPrefix(v, "-") + ` ` + sortType(v) + `,`
		}
		return sort[:len(sort)-1]
	}
	return ``
}

// Limit get limit
func (q *Query) Limit() int32 {
	return q.Count + 1
}

// Offset get Offset
func (q *Query) Offset() int32 {
	return (q.Page - 1) * q.Count
}

// Limit generate limit and offset for pagination
func (q *Query) LimitOffset() string {
	l := strconv.Itoa(int(q.Count + 1))
	o := strconv.Itoa((int(q.Page - 1)) * int(q.Count))
	return ` LIMIT ` + l + ` OFFSET ` + o
}

func translateOperator(s string) string {
	operator := Operator[strings.ToLower(s)]
	if operator == "" {
		return Operator["eq"]
	}
	return operator
}

// IsArgNil check type is null
func IsArgNil(i interface{}) bool {
	r := reflect.ValueOf(i)
	switch r.Kind() {
	case reflect.Slice:
		return r.Len() == 0
	case reflect.String:
		return r.String() == ""
	case reflect.Int:
		return r.Int() == 0
	case reflect.Int32:
		return r.Int() == 0
	case reflect.Int64:
		return r.Int() == 0
	case reflect.Float32:
		return r.Float() == 0
	case reflect.Float64:
		return r.Float() == 0
	default:
		return false
	}
}
