package farm

import (
	"database/sql"
	"fmt"

	"github.com/suradidchao/listenfield/repo"
)

// IAdapter is an interface for getting farm from db
type IAdapter interface {
	CreateFarm(farm repo.Farm, farmerID int) (farmID int, err error)
}

// MySQLAdapter is an farm adapter for getting Farm from MYSQL
type MySQLAdapter struct {
	table string
	db    *sql.DB
}

// CreateFarm is a method for inserting farm into mysql db
func (a MySQLAdapter) CreateFarm(farm repo.Farm, farmerID int) (farmID int, err error) {
	insertStatement := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?)", a.table)
	res, err := a.db.Exec(insertStatement, farm.FarmName, farm.CreatedDate, farmID)
	if err != nil {
		return farmID, err
	}
	lastID, err := res.LastInsertId()
	farmID = int(lastID)
	if err != nil {
		return farmID, err
	}
	return farmID, nil
}

// NewMySQLAdapter factory method for mysql adapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		db:    db,
		table: "farm",
	}
}
