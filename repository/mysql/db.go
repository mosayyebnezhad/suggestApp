package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDB struct {
	db *sql.DB
}

func New() *MySqlDB {

	MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD := "gameapp_db", "gameapp", "gameappt0lk2o20"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3308)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE))
	if err != nil {
		panic(fmt.Errorf("error when run app %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDB{
		db: db,
	}
}
