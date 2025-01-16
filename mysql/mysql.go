package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Global variable to hold the DB connection
var (
	DB       *sql.DB
	UserName = "admin"
	Password = "admin"
	Host     = "localhost"
	Port     = "3306"
	DBName   = "EH"
)

/*
 * leave the databaseName empty string "" then it will take Default DB name from DBName
 */
func ConnectToDatabase(databaseName string) (*sql.DB, error) {

	if databaseName == "" {
		databaseName = DBName
	}
	var err error
	// Replace with your actual MySQL connection details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		UserName,
		Password,
		Host,
		Port,
		databaseName,
	)
	fmt.Println("DSN:", dsn)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}

	// Check if the connection is successful
	err = DB.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
	return DB, nil
}

// Function to close the database connection
func CloseDatabaseConnection() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	fmt.Println("Database connection closed.")
}

// Function to fetch and scan the result into variadic parameters
func RunQuery(query string, args ...any) (*sql.Rows, error) {

	err := DB.Ping()
	if err != nil {
		DB, err = ConnectToDatabase("")
		if err != nil {
			return nil, fmt.Errorf("could not ping the database: %v", err)
		}
	}

	// Execute the query
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	// defer rows.Close() - clossing the row here making issue becuase if I am leaving the function the row data got vanished

	return rows, nil
}
