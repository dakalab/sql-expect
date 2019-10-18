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

type Model struct {
	ID    int `db:"id"`
	Value int `db:"value"`
	mock  sqlmock.Sqlmock
}

func (m Model) TableName() string {
	return "some_table"
}

func (m Model) Columns() []string {
	return []string{
		"id",
		"value",
	}
}

func (m Model) Mock() sqlmock.Sqlmock {
	return m.mock
}

type ModelTestSuite struct {
	suite.Suite
	db    *sql.DB
	model *Model
}

// TestModelTestSuite runs ModelTestSuite
func TestModelTestSuite(t *testing.T) {
	suite.Run(t, new(ModelTestSuite))
}

// SetupSuite runs once at the very start of the testing suite, before any tests are run.
func (suite *ModelTestSuite) SetupSuite() {
	log.Println("Model Tests Begin")

	suite.model = new(Model)
	suite.db, suite.model.mock, _ = sqlmock.New()
}

// TearDownSuite runs once at the very end of the testing suite, after all tests have been run.
func (suite *ModelTestSuite) TearDownSuite() {
	log.Println("Model Tests End")
}

func (suite *ModelTestSuite) TestUpdateModel() {
	UpdateModel(suite.model, nil, 1)
	res, err := suite.db.Exec("UPDATE " + suite.model.TableName() + " SET value = 2 WHERE id = 1")
	suite.NoError(err)
	rowsAffected, err := res.RowsAffected()
	suite.NoError(err)
	suite.Equal(int64(1), rowsAffected)

	UpdateModel(suite.model, errors.New("db error"), 0)
	_, err = suite.db.Exec("UPDATE " + suite.model.TableName() + " SET value = 2 WHERE id = 1")
	suite.Error(err)
}

func (suite *ModelTestSuite) TestSelect() {
	SelectModel(suite.model, errors.New("db error"))
	_, err := suite.db.Query("SELECT * FROM " + suite.model.TableName())
	suite.Error(err)

	SelectModel(suite.model, nil)
	rows, err := suite.db.Query("SELECT * FROM " + suite.model.TableName())
	suite.NoError(err)
	defer rows.Close()
	suite.False(rows.Next())

	SelectModel(suite.model, nil, []driver.Value{1, 2})
	rows, err = suite.db.Query("SELECT * FROM " + suite.model.TableName())
	suite.NoError(err)
	var id, value int
	for rows.Next() {
		err := rows.Scan(&id, &value)
		suite.NoError(err)
		suite.Equal(1, id)
		suite.Equal(2, value)
	}
}

func (suite *ModelTestSuite) TestInsert() {
	InsertModel(suite.model, nil, 1)
	res, err := suite.db.Exec("INSERT INTO " + suite.model.TableName() + "(id, value) VALUES (1, 2)")
	suite.NoError(err)
	lastInsertID, err := res.LastInsertId()
	suite.NoError(err)
	suite.Equal(int64(1), lastInsertID)

	InsertModel(suite.model, errors.New("db error"), 0)
	_, err = suite.db.Exec("INSERT INTO " + suite.model.TableName() + "(id, value) VALUES (1, 2)")
	suite.Error(err)
}

func (suite *ModelTestSuite) TestReplace() {
	ReplaceModel(suite.model, nil, 1)
	res, err := suite.db.Exec("REPLACE INTO " + suite.model.TableName() + "(id, value) VALUES (1, 2)")
	suite.NoError(err)
	lastInsertID, err := res.LastInsertId()
	suite.NoError(err)
	suite.Equal(int64(1), lastInsertID)

	ReplaceModel(suite.model, errors.New("db error"), 0)
	_, err = suite.db.Exec("REPLACE INTO " + suite.model.TableName() + "(id, value) VALUES (1, 2)")
	suite.Error(err)
}

func (suite *ModelTestSuite) TestDelete() {
	DeleteModel(suite.model, nil, 1)
	res, err := suite.db.Exec("DELETE FROM " + suite.model.TableName() + " WHERE id = 1")
	suite.NoError(err)
	rowsAffected, err := res.RowsAffected()
	suite.NoError(err)
	suite.Equal(int64(1), rowsAffected)

	DeleteModel(suite.model, errors.New("db error"), 0)
	_, err = suite.db.Exec("DELETE FROM " + suite.model.TableName() + " WHERE id = 1")
	suite.Error(err)
}

func (suite *ModelTestSuite) TestCount() {
	CountModel(suite.model, nil, 1)
	rows, err := suite.db.Query("SELECT * FROM " + suite.model.TableName())
	suite.NoError(err)
	defer rows.Close()
	var count int
	rows.Next()
	err = rows.Scan(&count)
	suite.NoError(err)
	suite.Equal(1, count)

	CountModel(suite.model, errors.New("db error"), 0)
	_, err = suite.db.Query("SELECT * FROM " + suite.model.TableName())
	suite.Error(err)
}
