package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgreDB() *sqlx.DB {
	db, err := connectPostgre()
	if err != nil {
		log.Fatal(err)
	}
	createTable(db)

	return db
}

func connectPostgre() (*sqlx.DB, error) {

	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	ssl := "disable"

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbName, ssl)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTable(db *sqlx.DB) {
	productTableQ := `
	DROP TABLE IF EXISTS tasks;
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		taskname VARCHAR(255),
		description TEXT,
		duedate DATE,
		priority VARCHAR(50),
		status VARCHAR(50),
		createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

	_, err := db.Exec(productTableQ)
	if err != nil {
		log.Panicln("error create table", err.Error())
	}
}
