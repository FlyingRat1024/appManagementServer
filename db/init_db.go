package db

import (
	"androidappServer/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBWork struct {
	Dsn string
}

func GetDB() (*sql.DB, error) {
	dbConfig := config.InitConfig().DB
	dbw := DBWork{
		Dsn: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DB),
	}
	db, err := sql.Open("mysql", dbw.Dsn)
	return db, err
}
