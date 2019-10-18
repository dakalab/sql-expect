package expect

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type ExpectTestSuite struct {
	suite.Suite
	db      *sql.DB
	mock    sqlmock.Sqlmock
	table   string
	columns []string
}

// TestExpectTestSuite runs ExpectTestSuite
func TestExpectTestSuite(t *testing.T) {
	suite.Run(t, new(ExpectTestSuite))
}

// SetupSuite runs once at the very start of the testing suite, before any tests are run.
func (suite *ExpectTestSuite) SetupSuite() {
	log.Println("Expect Tests Begin")

	suite.db, suite.mock, _ = sqlmock.New()
	suite.table = "some_table"
	suite.columns = []string{
		"id",
		"value",
	}
}

// TearDownSuite runs once at the very end of the testing suite, after all tests have been run.
func (suite *ExpectTestSuite) TearDownSuite() {
	log.Println("Expect Tests End")
}

func (suite *ExpectTestSuite) TestSelect() {
	Select(suite.mock, suite.table, suite.columns, errors.New("db error"))
	_, err := suite.db.Query("SELECT * FROM " + suite.table)
	suite.Error(err)

	Select(suite.mock, suite.table, suite.columns, nil)
	rows, err := suite.db.Query("SELECT * FROM " + suite.table)
	suite.NoError(err)
	defer rows.Close()
	suite.False(rows.Next())

	Select(suite.mock, suite.table, suite.columns, nil, []driver.Value{1, 2})
	rows, err = suite.db.Query("SELECT * FROM " + suite.table)
	suite.NoError(err)
	var id, value int
	for rows.Next() {
		err := rows.Scan(&id, &value)
		suite.NoError(err)
		suite.Equal(1, id)
		suite.Equal(2, value)
	}
}

func (suite *ExpectTestSuite) TestUpdate() {
	Update(suite.mock, suite.table, nil, 1)
	res, err := suite.db.Exec("UPDATE " + suite.table + " SET value = 2 WHERE id = 1")
	suite.NoError(err)
	rowsAffected, err := res.RowsAffected()
	suite.NoError(err)
	suite.Equal(int64(1), rowsAffected)

	Update(suite.mock, suite.table, errors.New("db error"), 0)
	_, err = suite.db.Exec("UPDATE " + suite.table + " SET value = 2 WHERE id = 1")
	suite.Error(err)
}

func (suite *ExpectTestSuite) TestInsert() {
	Insert(suite.mock, suite.table, nil, 1)
	res, err := suite.db.Exec("INSERT INTO " + suite.table + "(id, value) VALUES (1, 2)")
	suite.NoError(err)
	lastInsertID, err := res.LastInsertId()
	suite.NoError(err)
	suite.Equal(int64(1), lastInsertID)

	Insert(suite.mock, suite.table, errors.New("db error"), 0)
	_, err = suite.db.Exec("INSERT INTO " + suite.table + "(id, value) VALUES (1, 2)")
	suite.Error(err)
}

func (suite *ExpectTestSuite) TestReplace() {
	Replace(suite.mock, suite.table, nil, 1)
	res, err := suite.db.Exec("REPLACE INTO " + suite.table + "(id, value) VALUES (1, 2)")
	suite.NoError(err)
	lastInsertID, err := res.LastInsertId()
	suite.NoError(err)
	suite.Equal(int64(1), lastInsertID)

	Replace(suite.mock, suite.table, errors.New("db error"), 0)
	_, err = suite.db.Exec("REPLACE INTO " + suite.table + "(id, value) VALUES (1, 2)")
	suite.Error(err)
}

func (suite *ExpectTestSuite) TestDelete() {
	Delete(suite.mock, suite.table, nil, 1)
	res, err := suite.db.Exec("DELETE FROM " + suite.table + " WHERE id = 1")
	suite.NoError(err)
	rowsAffected, err := res.RowsAffected()
	suite.NoError(err)
	suite.Equal(int64(1), rowsAffected)

	Delete(suite.mock, suite.table, errors.New("db error"), 0)
	_, err = suite.db.Exec("DELETE FROM " + suite.table + " WHERE id = 1")
	suite.Error(err)
}

func (suite *ExpectTestSuite) TestCount() {
	Count(suite.mock, suite.table, nil, 1)
	rows, err := suite.db.Query("SELECT * FROM " + suite.table)
	suite.NoError(err)
	defer rows.Close()
	var count int
	rows.Next()
	err = rows.Scan(&count)
	suite.NoError(err)
	suite.Equal(1, count)

	Count(suite.mock, suite.table, errors.New("db error"), 0)
	_, err = suite.db.Query("SELECT * FROM " + suite.table)
	suite.Error(err)
}
