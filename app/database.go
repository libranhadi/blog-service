package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func NewDb() (*sql.DB, error) {
	// Create the database if it doesn't exist
	if err := createDatabaseIfNotExists(); err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/articles")
	if err != nil {
		return nil, err
	}

	if err := createTableIfNotExists(db); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func createDatabaseIfNotExists() error {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS articles")
	if err != nil {
		panic(err)
	}

	return nil
}

func createTableIfNotExists(db *sql.DB) error {

	sqlStatements, err := os.ReadFile("./app/post_table.sql")
	if err != nil {
		return err
	}

	sqlString := string(sqlStatements)

	_, err = db.Exec(sqlString)
	if err != nil {
		return fmt.Errorf("error executing SQL statements: %w", err)
	}

	return nil
}
