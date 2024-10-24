package main

import (
	"database/sql"

	"github.com/labstack/gommon/log"
)

// Ideally this would be migrations, but for now this should be fine
// https://gorm.io/docs/migration.html
// TODO: Use gorm migrator instead of raw sql

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

	insertTestData(db)
}

func insertTestData(db *sql.DB) {
	insertDataStmt := `
		INSERT INTO Books
		VALUES 
			( '9783641279110', '1984', 'George Orwell', 'June 8, 1949', 1),
			( '9780198826736', 'Pride and Prejudice', 'Jane Austen', 'January 28, 1813', 0),
			( '9780007322596', 'The Lord of the Rings', 'J. R. R. Tolkien', 'July 29, 1954', 1),
			( '9781438119250', 'The Catcher in the Rye', 'J. D. Salinger', 'July 16, 1951', 0)
	`

	_, err := db.Exec(insertDataStmt)
	if err != nil {
		log.Errorf("%q: %s\n", err, insertDataStmt)
		return
	}
}
