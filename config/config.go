package config

import (
	"suggestApp/repository/mysql"
	"suggestApp/service/authservice"
)

type HttpServer struct {
	Port int
}

type Config struct {
	HttpServer HttpServer
	Auth       authservice.Config
	Mysql      mysql.Config
}
