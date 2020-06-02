package qibo

const (
	defaultPage  = 1
	defaultCount = 10
)

// NewQuery clone to a new query builder with initial value
func NewQuery(sort string, filter map[string]interface{}) *Query {
	query := &Query{
		Sort:   sort,
		Filter: filter,
	}

	return query
}

// NewPagination initiate new pagination obj
func NewPagination(page int32, count int32) *Pagination {
	var p, c int32
	if page == 0 {
		p = defaultPage
	} else if page > 0 {
		p = page
	}

	if count == 0 {
		c = defaultCount
	} else if count > 0 {
		c = count
	}else if count < 0{
		c = count
	}

	pagination := &Pagination{
		CurrentPage: p,
		PageSize:    c,
		TotalResult: 0,
		TotalPage:   0,
	}

	return pagination
}
