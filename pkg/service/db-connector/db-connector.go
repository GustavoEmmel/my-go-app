package db_connector

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"my-go-app/config/database"
)

func Conn() (db *sql.DB) {
	dbDriver := database.MYSQL_DRIVER
	dbUser := database.MYSQL_USER
	dbPass := database.MYSQL_PASS
	dbName := database.MYSQL_DBNAME
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
