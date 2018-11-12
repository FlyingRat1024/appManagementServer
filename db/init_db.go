package db

import (
	"database/sql"
	"fmt"
)
import (
	"androidappServer/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	dbConfig := config.InitConfig().DB
	dbfmt := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DB)
	db, err := sql.Open("mysql", dbfmt)
	return db, err
}
