package db

import (
	"database/sql"
	"fmt"
	"log"

	"packages.vrianta.in/golang/db/mysql"
	"packages.vrianta.in/golang/db/util"
)

var (
	DB       *sql.DB
	UserName = "admin"
	Password = "admin"
	Host     = "localhost"
	Port     = "3306"
	DBName   = "EH"

	DBType = DBTypes.MYSQL

	DBTypes = struct {
		MYSQL int
	}{
		MYSQL: 0,
	}
)

func Init() {
	if DBType == DBTypes.MYSQL {
		mysql.DB = DB
		mysql.UserName = util.GetEnvriontmentVar("MYSQL_USERNAME", UserName)
		mysql.Password = util.GetEnvriontmentVar("MYSQL_PASSWORD", Password)
		mysql.Host = util.GetEnvriontmentVar("MYSQL_HOST", Host)
		mysql.Port = util.GetEnvriontmentVar("MYSQL_PORT", Port)
		mysql.DBName = util.GetEnvriontmentVar("MYSQL_DB_NAME", DBName)
	}
}

func ConnectToDatabase(databaseName string) (*sql.DB, error) {

	if databaseName == "" {
		databaseName = DBName
	}

	switch DBType {
	case DBTypes.MYSQL:
		return mysql.ConnectToDatabase(databaseName)
	default:
		return nil, fmt.Errorf("no proper dbType is set")
	}

}

func CloseDatabaseConnection() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	fmt.Println("Database connection closed.")
}

func RunQuery(query string, args ...any) (*sql.Rows, error) {
	switch DBType {
	case DBTypes.MYSQL:
		return mysql.RunQuery(query)
	default:
		return nil, fmt.Errorf("no proper dbType is set")
	}

}
