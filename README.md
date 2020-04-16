# Qibo

[![Build Status](https://travis-ci.com/qasir-id/qibo.svg?branch=master)](https://travis-ci.com/qasir-id/qibo)

Qibo is query builder for [Go](https://golang.org/) which is used internally by [Qasir](https://www.qasir.id/) Tech.

## Instalation

```
# Go modules
$> go get -u github.com/qasir-id/qibo
$> go mod tidy
```

## String Query Operators
| Code   | Sql operator | Description |
| ------ | ------ | ------ |
| "gt" |   ">" | greater than |
| "lt" |   "<" | lower then |
| "eq" |   "=" | equal |
| "ne" |   "!=" | not equal |
| "gte" |  ">=" | greater than equal |
| "lte" |  "<=" | lower then |
| "like" | "LIKE" | like / contains |
| "in" |   "IN" | array |

## Sort
| Code   | Sql operator | Description |
| ------ | ------ | ------ |
| "-" |   "DESC" | Descending |
| " " |   "ASC" | Ascending |


use 

## Basic Usage

```go
import "github.com/qasir-id/qibo"
import "github.com/jinzhu/gorm"


query := qibo.NewQuery(0, 0, "-name" { // sort by name descending
	"id$in!":  	        []int{23, 25}, // mandatory filter
	"name$like":        "Sample name",
	"description$like": "Sample description",
	"created_at$gte":   "2019-12-01",
	"created_at$lte":   "2019-12-31",
})

smt, args := query.Where()
var categories []model.Category

res := db.Where(stmt, args...).Find(&categories).Order(query.Order())
```
