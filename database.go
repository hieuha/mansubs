package main

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DatabaseSource string
}

func (d Database) connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", d.DatabaseSource)
	return db, err
}

func (d Database) isTableTargetExist(db *sql.DB) bool {
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name='targets';`
	var tableName string

	err := db.QueryRow(query).Scan(&tableName)
	if err != nil {
		return false
	}
	if tableName != "" {
		return true
	}
	return false
}

func (d Database) createTableTargets(db *sql.DB) error {
	createTargetsTableSQL := `CREATE TABLE "targets" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"domain" VARCHAR(255) NULL,
		"subdomain" VARCHAR(255) NULL,
		"technology" VARCHAR(255) NULL,
		"created" DATETIME
	);
	`
	statement, err := db.Prepare(createTargetsTableSQL)

	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}

func (d Database) addTarget(db *sql.DB, target Target) error {
	query := `INSERT INTO targets("domain", "subdomain", "technology", "created") 
			  VALUES (?, ?, ?, CURRENT_TIMESTAMP);
	`
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(target.Domain, target.Subdomain, target.Technology)
	if err == nil {
		fmt.Println("Added", target)
	}
	return err
}
