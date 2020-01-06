package qibo_test

import (
	"testing"

	"github.com/qasir-id/qibo"
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
	t.Pagination = qibo.NewPagination(1, 10)
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

func (t *Pagination) TestLimitOffset() {
	loExpected := "LIMIT 10 OFFSET 0"
	loActual := t.Pagination.LimitOffset()

	t.Require().Equal(loExpected, loActual)
}
