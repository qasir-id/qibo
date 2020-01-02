package qibo_test

import (
	"log"
	"testing"

	"github.com/QasirID/qibo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type Query struct {
	suite.Suite
	mock  mock.Mock
	query *qibo.Query
}

func TestQuery(t *testing.T) {
	suite.Run(t, new(Query))
}

func (t *Query) SetupTest() {
	t.query = qibo.NewQuery(0, 0, "", map[string]interface{}{})
}

// func (t *Query) AfterTest(_, _ string) {
// 	require.NoError(t.T(), t.mock.ExpectationsWereMet())
// }

func (t *Query) TestQueryUser() {
	t.query.SetFilter(map[string]interface{}{
		"user_id$eq": 1,
	})
	smt, args := t.query.Where()
	t.Require().Equal(smt, " user_id = ? ")
	t.Require().Equal(args, []interface{}{1})
	log.Println("Statement: ", smt, "Arguments: ", args)
}
