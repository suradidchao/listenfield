package field

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/entity"
)

// IAdapter is an interface of field interacting with field in db
type IAdapter interface {
	Create(field entity.Field) (fieldID int, err error)
}

// MySQLAdapter is a mysql adapter of field
type MySQLAdapter struct {
	db    *sql.DB
	table string
}

// Create is a method for inserting field to farm in field table
func (a MySQLAdapter) Create(field entity.Field) (fieldID int, err error) {
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?, ?, ?, ?)", a.table)
	res, err := a.db.Exec(insertStmt, field.FarmID, field.Crop, field.Status, field.Area, time.Now(), field.FieldName)
	if err != nil {
		return fieldID, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return fieldID, err
	}
	fieldID = int(lastInsertID)
	return fieldID, nil
}

// NewMySQLAdapter is a factory method for field mysql adapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		db:    db,
		table: "field",
	}
}
