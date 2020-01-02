package qibo

const (
	defaultPage  = 1
	defaultCount = 10
)

// NewQuery clone to a new query builder with initial value
func NewQuery(page int32, count int32, sort string, filter map[string]interface{}) *Query {
	var p, c int32
	if page == 0 {
		p = defaultPage
	}
	if count == 0 {
		c = defaultCount
	}
	query := &Query{
		Page:   p,
		Count:  c,
		Sort:   sort,
		Filter: filter,
	}

	return query
}

// NewPagination initiate new pagination obj
func NewPagination(q *Query, totalResult int32) *Pagination {
	pageSize := totalResult
	currentPage := q.Page
	if currentPage == 0 {
		currentPage = defaultPage
	}
	if q.Count == 0 {
		pageSize = defaultCount
	} else if q.Count > 0 {
		pageSize = q.Count
	}

	p := &Pagination{
		CurrentPage: currentPage,
		PageSize:    pageSize,
		TotalResult: totalResult,
	}
	return p.SetTotalPage(totalResult)
}
