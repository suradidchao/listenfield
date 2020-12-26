package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/suradidchao/listenfield/entity"
	"golang.org/x/crypto/bcrypt"
)

// IAdapter is an interface for user adapter
type IAdapter interface {
	Create(user entity.User) (uid int, err error)
	hashPassword(password []byte) (pwd string, err error)
}

// MySQLAdapter is an user adapter for managing Uarm from MYSQL
type MySQLAdapter struct {
	table string
	db    *sql.DB
}

// Create is an adapter method for creating user in mysql
func (a MySQLAdapter) Create(user entity.User) (uid int, err error) {
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, ?, ?, ?)", a.table)
	hashedPwd, err := a.hashPassword([]byte(user.Password))
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

func (a MySQLAdapter) hashPassword(password []byte) (pwd string, err error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return pwd, err
	}
	return string(hash), nil
}

// NewMySQLAdapter is a factory method for user mysqlAdapter
func NewMySQLAdapter(db *sql.DB) MySQLAdapter {
	return MySQLAdapter{
		table: "user",
		db:    db,
	}
}
