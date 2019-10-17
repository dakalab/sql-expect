package expect

import (
	"database/sql/driver"
	"fmt"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// Select expects db select
func Select(mock sqlmock.Sqlmock, table string, columns []string, err error, values ...[]driver.Value) {
	sql := fmt.Sprintf("SELECT (.+) FROM %s", table)

	if err != nil {
		mock.ExpectQuery(sql).WillReturnError(err)
		return
	}

	if values == nil || len(values) == 0 {
		mock.ExpectQuery(sql).WillReturnRows(&sqlmock.Rows{})
		return
	}

	rows := sqlmock.NewRows(columns)
	for _, value := range values {
		rows.AddRow(value...)
	}

	mock.ExpectQuery(sql).WillReturnRows(rows)
}

// Update expects db update
func Update(mock sqlmock.Sqlmock, table string, err error, rowsAffected int64) {
	sql := fmt.Sprintf("UPDATE %s SET", table)

	if err != nil {
		mock.ExpectExec(sql).WillReturnError(err)
		return
	}

	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, rowsAffected))
}

// Insert expects db insert
func Insert(mock sqlmock.Sqlmock, table string, err error, lastInsertID int64) {
	sql := fmt.Sprintf("INSERT INTO %s", table)

	if err != nil {
		mock.ExpectExec(sql).WillReturnError(err)
		return
	}

	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(lastInsertID, 0))
}

// Replace expects db replace
func Replace(mock sqlmock.Sqlmock, table string, err error, lastInsertID int64) {
	sql := fmt.Sprintf("REPLACE INTO %s", table)

	if err != nil {
		mock.ExpectExec(sql).WillReturnError(err)
		return
	}

	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(lastInsertID, 0))
}

// Delete expects db delete
func Delete(mock sqlmock.Sqlmock, table string, err error, rowsAffected int64) {
	sql := fmt.Sprintf("DELETE FROM %s", table)

	if err != nil {
		mock.ExpectExec(sql).WillReturnError(err)
		return
	}

	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, rowsAffected))
}
