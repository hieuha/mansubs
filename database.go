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
		"created" DATETIME,
		unique (domain, subdomain)
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

func (d Database) getTargets(db *sql.DB, domain string) ([]Target, error) {
	var targets []Target
	query := `SELECT "id", "domain", "subdomain", "technology" FROM "targets" WHERE "domain" = ? ORDER BY "id"`
	statement, err := db.Prepare(query)
	if err != nil {
		return targets, err
	}

	rows, err := statement.Query(domain)

	if err != nil {
		return targets, err
	}

	defer rows.Close()
	for rows.Next() {
		var target = Target{}
		rows.Scan(&target.Id, &target.Domain, &target.Subdomain, &target.Technology)
		targets = append(targets, target)
	}

	return targets, nil
}

func (d Database) cleanTech(db *sql.DB) {
	queryClean := `UPDATE "targets" SET "technology" = ""`
	statementClean, err := db.Prepare(queryClean)
	_, err = statementClean.Exec()
	if err != nil {
		fmt.Println(err)
	}
}

func (d Database) updateTech(db *sql.DB, subdomain string, technology string) error {
	query := `UPDATE "targets" SET "technology" = technology || ? WHERE "subdomain" = ?`
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(fmt.Sprintf("%s,", technology), subdomain)
	if err != nil {
		return err
	}
	return nil
}

func (d Database) searchTargetByTech(db *sql.DB, tech string) ([]Target, error) {
	var targets []Target
	query := `SELECT "id", "domain", "subdomain", "technology" FROM "targets" WHERE "technology" LIKE ? ORDER BY "id"`

	statement, err := db.Prepare(query)
	if err != nil {
		return targets, err
	}
	tech = fmt.Sprintf("%%%s%%", tech)
	rows, err := statement.Query(tech)
	if err != nil {
		return targets, err
	}

	defer rows.Close()
	for rows.Next() {
		var target = Target{}
		rows.Scan(&target.Id, &target.Domain, &target.Subdomain, &target.Technology)
		targets = append(targets, target)
	}

	return targets, nil
}
