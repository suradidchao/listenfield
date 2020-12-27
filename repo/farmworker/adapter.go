package farmworker

import (
	"database/sql"
	"fmt"
)

// IAdapter is an interface of farmworker interacting with farmworker in db
type IAdapter interface {
	Create(farmID int, userID int) (fwID int, err error)
	GetAllByFarmID(farmID int) (userIDs []int, err error)
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

// GetAllByFarmID is a method for getting all farm workers of a farm
func (a MySQLAdapter) GetAllByFarmID(farmID int) (userIDs []int, err error) {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE farm_id=?", a.table)
	rows, err := a.db.Query(query, farmID)
	if err != nil {
		return userIDs, err
	}
	userIDs = []int{}
	for rows.Next() {
		var userID int
		switch err = rows.Scan(&userID); err {
		case sql.ErrNoRows:
			return userIDs, err
		case nil:
			userIDs = append(userIDs, userID)
		default:
			return userIDs, err
		}
	}
	return userIDs, nil
}

// NewMySQLAdapter is a factory method for farmworker mysql adapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		db:    db,
		table: "farm_worker",
	}
}
