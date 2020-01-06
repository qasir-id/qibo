package qibo_test

import (
	"strings"
	"testing"

	"github.com/qasir-id/qibo"
	"github.com/stretchr/testify/suite"
)

// Query struct for testing
type Query struct {
	suite.Suite
	// mock  mock.Mock
	query *qibo.Query
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(Query))
}

func (t *Query) SetupTest() {
	t.query = qibo.NewQuery("-name", map[string]interface{}{
		"user_id$eq":     1,
		"name$like":      "sample user",
		"created_at$gte": "2018-12-01",
		"created_at$lte": "2018-12-31",
	})
}

// func (t *Query) AfterTest(_, _ string) {
// 	require.NoError(t.T(), t.mock.ExpectationsWereMet())
// }

func (t *Query) TestFilter() {
	smt, args := t.query.Where()
	wheresExpected := []string{"user_id = ?", "name LIKE ?", "created_at >= ?", "created_at <= ?"}
	wheresActual := strings.Split(smt, " AND ")
	t.Require().Equal(wheresActual, wheresExpected)
	t.Require().Equal(args, []interface{}{1, "%sample user%", "2018-12-01 00:00:00", "2018-12-31 23:59:59"})
}

func (t *Query) TestOrder() {
	orderExpected := "name DESC"
	orderActual := t.query.Order()

	t.Require().Equal(orderExpected, orderActual)
}
