package expect

import (
	"database/sql/driver"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

// Modeler is the interface of DB Model
type Modeler interface {
	TableName() string
	Columns() []string
	Mock() sqlmock.Sqlmock
}

// SelectModel expects select from model
func SelectModel(m Modeler, err error, values ...[]driver.Value) {
	Select(m.Mock(), m.TableName(), m.Columns(), err, values...)
}

// UpdateModel expects update to model
func UpdateModel(m Modeler, err error, rowsAffected int64) {
	Update(m.Mock(), m.TableName(), err, rowsAffected)
}

// InsertModel expects insert into model
func InsertModel(m Modeler, err error, lastInsertID int64) {
	Insert(m.Mock(), m.TableName(), err, lastInsertID)
}

// ReplaceModel expects replace into model
func ReplaceModel(m Modeler, err error, lastInsertID int64) {
	Replace(m.Mock(), m.TableName(), err, lastInsertID)
}

// DeleteModel expects delete model
func DeleteModel(m Modeler, err error, rowsAffected int64) {
	Delete(m.Mock(), m.TableName(), err, rowsAffected)
}

// CountModel expects Count model
func CountModel(m Modeler, err error, count uint32) {
	Count(m.Mock(), m.TableName(), err, count)
}
