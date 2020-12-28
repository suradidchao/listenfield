package activity

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/entity"
)

// IAdapter is an interface for activity adapter
type IAdapter interface {
	Create(actvity entity.Activity) (aid int, err error)
}

// MySQLAdapter is an activity adapter for managing activity from MYSQL
type MySQLAdapter struct {
	table string
	db    *sql.DB
}

// Create is an adapter method for creating activity in mysql
func (a MySQLAdapter) Create(activity entity.Activity) (aid int, err error) {
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?, ?, ?, ?, ?, ?, ?)", a.table)
	res, err := a.db.Exec(insertStmt, activity.FarmID, activity.Field.FieldID, activity.Tractor.TractorID, activity.User.UserID, activity.ActivityName, activity.Area, activity.Cost, activity.Revenue, time.Now())
	if err != nil {
		return aid, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return aid, err
	}
	aid = int(lastInsertID)
	return aid, nil
}

// NewMySQLAdapter is a factory method for activity mysqlAdapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		table: "activity",
		db:    db,
	}
}
