package database

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

func New(dbPath string) *sql.DB {

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		log.Fatal("DB nil!")
	}

	migrate(db)

	return db

}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("DB migration error: %s",err)
	}
}
