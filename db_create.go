package main

import (
	"database/sql"

	"github.com/labstack/gommon/log"
)

func createBooksTable(db *sql.DB) {
	tableExistsStmt := `SELECT count(name) FROM sqlite_master WHERE type='table' AND name='Books';`

	row:= db.QueryRow(tableExistsStmt)
	var exists int 
	row.Scan(&exists)

	if (exists == 1) { 
		log.Print("Table alreay exists, skipping creation")
		return 
	}

	createStmt := `
		CREATE TABLE Books (
			ISBN nvarchar(17) not null,
			Title nvarchar(255) not null,
			Author nvarchar(100) null,
			ReleaseDate date null,
			Availability bool not null
		);
	`

	_, err := db.Exec(createStmt)
	if err != nil {
		log.Errorf("%q: %s\n", err, createStmt)
		return
	}
}
