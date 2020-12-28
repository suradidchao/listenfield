package tractor

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/entity"
)

// IAdapter is an interface for getting farm from db
type IAdapter interface {
	Create(tractor entity.Tractor) (tractorID int, err error)
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

// NewMySQLAdapter is a factory method for tractor MySQLAdapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		table: "tractor",
		db:    db,
	}

}
