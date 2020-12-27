package farmworker

import (
	"database/sql"
	"fmt"
)

// IAdapter is an interface of farmworker interacting with farmworker in db
type IAdapter interface {
	Create(farmID int, userID int) (fwID int, err error)
}

// MySQLAdapter is a mysql adapter of farmworker
type MySQLAdapter struct {
	db    *sql.DB
	table string
}

// Create is a method for inserting user to farm in farm worker table
func (a MySQLAdapter) Create(farmID int, userID int) (fwID int, err error) {
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?)", a.table)
	res, err := a.db.Exec(insertStmt, farmID, userID)
	if err != nil {
		return fwID, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return fwID, err
	}
	fwID = int(lastInsertID)
	return fwID, nil
}

// NewMySQLAdapter is a factory method for farmworker mysql adapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		db:    db,
		table: "farm_worker",
	}
}
