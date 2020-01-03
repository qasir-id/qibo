package qibo_test

import (
	"testing"

	"github.com/QasirID/qibo"
	"github.com/stretchr/testify/suite"
)

// Pagination struct for testing
type Pagination struct {
	suite.Suite
	// mock  mock.Mock
	Pagination *qibo.Pagination
}

func TestPagination(t *testing.T) {
	suite.Run(t, new(Pagination))
}

func (t *Pagination) SetupTest() {
	query := qibo.NewQuery(0, 0, "", map[string]interface{}{
		"user_id$eq": 1,
	})
	t.Pagination = qibo.NewPagination(query, 32)
}

func (t *Pagination) TestTotalPage() {
	pExpected := &qibo.Pagination{
		CurrentPage: 1,
		PageSize:    10,
		TotalPage:   24,
		TotalResult: 240,
	}
	pActual := t.Pagination.SetTotalPage(240)
	t.Require().Equal(pExpected, pActual)
	// log.Println("Statement: ", smt, "Arguments: ", args)
}
