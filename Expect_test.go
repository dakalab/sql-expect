package expect

import (
	"database/sql/driver"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestUpdate(t *testing.T) {
	_, mock, _ := sqlmock.New()
	Update(mock, "some_table", nil, 1)
	Update(mock, "some_table", errors.New("db error"), 0)
}

func TestSelect(t *testing.T) {
	_, mock, _ := sqlmock.New()
	table := "some_table"
	columns := []string{
		"id",
		"value",
	}
	Select(mock, table, columns, errors.New("db error"))
	Select(mock, table, columns, nil)
	Select(mock, table, columns, nil, []driver.Value{1, 2})
}

func TestInsert(t *testing.T) {
	_, mock, _ := sqlmock.New()
	Insert(mock, "some_table", nil, 1)
	Insert(mock, "some_table", errors.New("db error"), 0)
}

func TestReplace(t *testing.T) {
	_, mock, _ := sqlmock.New()
	Replace(mock, "some_table", nil, 1)
	Replace(mock, "some_table", errors.New("db error"), 0)
}

func TestDelete(t *testing.T) {
	_, mock, _ := sqlmock.New()
	Delete(mock, "some_table", nil, 1)
	Delete(mock, "some_table", errors.New("db error"), 0)
}
