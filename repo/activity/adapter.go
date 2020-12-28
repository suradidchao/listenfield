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
	AggCostAndRevenueByFarmID(farmID int, startDate time.Time, endDate time.Time) (cr CostAndRevenue, err error)
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

// AggCostAndRevenueByFarmID is an adapter method for aggregate cost summary from mysql
func (a MySQLAdapter) AggCostAndRevenueByFarmID(farmID int, startDate time.Time, endDate time.Time) (cr CostAndRevenue, err error) {
	query := fmt.Sprintf("SELECT SUM(cost), SUM(revenue) FROM %s WHERE created_date > ? AND created_date < ?", a.table)
	row := a.db.QueryRow(query, startDate.Format("2006-01-02 15:04:05"), endDate.Format("2006-01-02 15:04:05"))
	err = row.Scan(&cr.Cost, &cr.Revenue)
	if err != nil {
		return cr, err
	}
	return cr, nil
}

// NewMySQLAdapter is a factory method for activity mysqlAdapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		table: "activity",
		db:    db,
	}
}
