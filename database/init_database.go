package database

import (
	"androidappServer/config"
	"database/sql"
	"fmt"
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
