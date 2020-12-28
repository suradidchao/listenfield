package field

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/app/entity"
)

// IAdapter is an interface of field interacting with field in db
type IAdapter interface {
	Create(field entity.Field) (fieldID int, err error)
	Delete(fieldID int) (err error)
	Update(fieldID int, field entity.Field) (err error)
}

// MySQLAdapter is a mysql adapter of field
type MySQLAdapter struct {
	db    *sql.DB
	table string
}

// Create is a method for inserting field to farm in field table
func (a MySQLAdapter) Create(field entity.Field) (fieldID int, err error) {
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?, ?, ?)", a.table)
	res, err := a.db.Exec(insertStmt, field.FarmID, field.FieldName, field.Crop, field.Area, time.Now())
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

// Delete a field
func (a MySQLAdapter) Delete(fieldID int) (err error) {
	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", a.table)
	_, err = a.db.Exec(deleteStmt, fieldID)
	if err != nil {
		return err
	}
	return nil
}

// Update is a method for updating field values
func (a MySQLAdapter) Update(fieldID int, field entity.Field) (err error) {
	updateStmt := fmt.Sprintf("UPDATE %s SET field_name = ?, farm_id = ?, crop = ?, area = ? WHERE id = ?", a.table)
	_, err = a.db.Exec(updateStmt, field.FieldName, field.FarmID, field.Crop, field.Area, fieldID)
	if err != nil {
		return err
	}
	return nil
}

// NewMySQLAdapter is a factory method for field mysql adapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		db:    db,
		table: "field",
	}
}
