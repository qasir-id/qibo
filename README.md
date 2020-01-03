# Qibo

[![Build Status](https://travis-ci.com/QasirID/qibo.svg?branch=master)](https://travis-ci.com/QasirID/qibo)

Qibo is query builder for [Go](https://golang.org/) which is used internally by [Qasir](https://www.qasir.id/) Tech.

## Instalation

```
# Go modules
go mod tidy
```

## Basic Usage

```go
import "github.com/QasirID/qibo"
import "github.com/jinzhu/gorm"

query := qibo.NewQuery(0, 0, "name" {
	"id$in!":  	        []int{23, 25},
	"name$like":        "Sample name",
	"description$like": "Sample description",
	"created_at$gte":   "2019-12-01",
	"created_at$lte":   "2019-12-31",
})

smt, args := query.Where()
var categories []model.Category

res := db.Where(stmt, args...).Find(&categories).Order(query.Order())
```
