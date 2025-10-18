package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string
	Password string
	Port     int
	Host     string
	DbName   string
}

type MySqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MySqlDB {

	MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD, HOST, PORT := config.DbName, config.Username, config.Password, config.Host, config.Port
	// MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD, HOST, PORT := "gameapp_db", "gameapp", "gameappt0lk2o20", "localhost", 3306
	// MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD := "gameapp_db", "gameapp", "gameappt0lk2o20"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s", MYSQL_USER, MYSQL_PASSWORD, HOST, PORT, MYSQL_DATABASE))
	if err != nil {
		panic(fmt.Errorf("error when run app %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySqlDB{
		db:     db,
		config: config,
	}
}
