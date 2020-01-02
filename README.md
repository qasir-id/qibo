# Qibo
Qibo is query builder for [Go](https://golang.org/) which is used internally by [Qasir](https://www.qasir.id/) Tech.

## Instalattion
```
# Go modules
go mod tidy
```

## Basic Usage
```go
import "github.com/QasirID/qibo"

// query filter

query := qibo.Query{
	Filter: map[string]interface{}{
		"user_id$in!":  	123,
		"name$like":        "Sample name",
		"description$like": Sample description,
		"created_at$gte":   "2019-12-01",
		"created_at$lte":   "2019-12-31",
	},
	Page:  0,
	Count: 0, // default is 10
	Sort:  "name",
}
```	
