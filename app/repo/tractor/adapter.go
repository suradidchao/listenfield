package tractor

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/app/entity"
)

// IAdapter is an interface for getting farm from db
type IAdapter interface {
	Create(tractor entity.Tractor) (tractorID int, err error)
	Delete(tractorID int) (err error)
	Update(tractorID int, tractor entity.Tractor) (err error)
}

// MySQLAdapter is an tractor adapter for operating with tractor from MYSQL
type MySQLAdapter struct {
	table string
	db    *sql.DB
}

// Create is a method for inserting tractor into mysql db
func (a MySQLAdapter) Create(tractor entity.Tractor) (tractorID int, err error) {
	insertStatement := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?)", a.table)
	res, err := a.db.Exec(insertStatement, tractor.TractorName, tractor.TractorID, time.Now())
	if err != nil {
		return tractorID, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return tractorID, err
	}
	tractorID = int(lastID)
	return tractorID, nil
}

// Delete a tractor
func (a MySQLAdapter) Delete(tractorID int) (err error) {
	deleteStmt := fmt.Sprintf("DELETE FROM %s WHERE id=?", a.table)
	_, err = a.db.Exec(deleteStmt, tractorID)
	if err != nil {
		return err
	}
	return nil
}

// Update is a method for updating tractor attribute in a mysql db
func (a MySQLAdapter) Update(tractorID int, tractor entity.Tractor) (err error) {
	updateStmt := fmt.Sprintf("UPDATE %s SET tractor_name = ?, farm_id = ? WHERE id = ?", a.table)
	_, err = a.db.Exec(updateStmt, tractor.TractorName, tractor.FarmID, tractorID)
	if err != nil {
		return err
	}
	return nil
}

// NewMySQLAdapter is a factory method for tractor MySQLAdapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		table: "tractor",
		db:    db,
	}

}
