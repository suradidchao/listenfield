package farm

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/app/entity"
)

// IAdapter is an interface for getting farm from db
type IAdapter interface {
	CreateFarm(farm entity.Farm, farmerID int) (farmID int, err error)
	GetFarmIDsByUserID(userID int) (farmIDs []int, err error)
}

// MySQLAdapter is an farm adapter for getting Farm from MYSQL
type MySQLAdapter struct {
	table string
	db    *sql.DB
}

// CreateFarm is a method for inserting farm into mysql db
func (a MySQLAdapter) CreateFarm(farm entity.Farm, farmerID int) (farmID int, err error) {
	insertStatement := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?)", a.table)
	res, err := a.db.Exec(insertStatement, farm.FarmName, farmerID, time.Now())
	if err != nil {
		return farmID, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return farmID, err
	}
	farmID = int(lastID)
	return farmID, nil
}

// GetFarmIDsByUserID is a method for getting farm ids own by a user
func (a MySQLAdapter) GetFarmIDsByUserID(userID int) (farmIDs []int, err error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE farm_owner_id=?", a.table)
	rows, err := a.db.Query(query, userID)
	if err != nil {
		return farmIDs, err
	}
	farmIDs = []int{}
	for rows.Next() {
		var farmID int
		switch err = rows.Scan(&farmID); err {
		case sql.ErrNoRows:
			return farmIDs, err
		case nil:
			farmIDs = append(farmIDs, farmID)
		default:
			return farmIDs, err
		}
	}
	return farmIDs, nil
}

// NewMySQLAdapter factory method for mysql adapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		db:    db,
		table: "farm",
	}
}
