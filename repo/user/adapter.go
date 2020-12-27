package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/internal/passgen"
)

// IAdapter is an interface for user adapter
type IAdapter interface {
	Create(user entity.User) (uid int, err error)
	GetByUsername(username string) (user entity.User, err error)
}

// MySQLAdapter is an user adapter for managing Uarm from MYSQL
type MySQLAdapter struct {
	table string
	db    *sql.DB
}

// Create is an adapter method for creating user in mysql
func (a MySQLAdapter) Create(user entity.User) (uid int, err error) {
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?, ?)", a.table)
	hashedPwd, err := passgen.HashPassword([]byte(user.Password))
	if err != nil {
		return uid, err
	}
	res, err := a.db.Exec(insertStmt, user.Username, hashedPwd, user.Email, time.Now())
	if err != nil {
		return uid, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return uid, err
	}
	uid = int(lastInsertID)
	return uid, nil
}

// GetByUsername is a method for getting a user from db by username
func (a MySQLAdapter) GetByUsername(username string) (user entity.User, err error) {
	query := fmt.Sprintf("SELECT id, username, password, email, created_date FROM %s WHERE username=?", a.table)
	row := a.db.QueryRow(query, username)
	err = row.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedDate)
	if err != nil {
		return user, err
	}
	return user, nil
}

// NewMySQLAdapter is a factory method for user mysqlAdapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		table: "user",
		db:    db,
	}
}
