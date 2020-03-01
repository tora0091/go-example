package database

import (
	"database/sql"
	"log"

	"go-example/restfulapi/config"
)

func InitializeDatabase() {
	db := DBConnection()
	_, err := db.Exec(`create table if not exists users (
		"id" integer primary key,
		"name" text,
		"email" text,
		"job" text,
		"created_at" text not null default (datetime('now', 'localtime')),
		"updated_at" text not null default (datetime('now', 'localtime')),
		"deleted_at" text 
	)`)
	if err != nil {
		log.Fatalln(err)
	}
}

func DBConnection() *sql.DB {
	db, err := sql.Open(config.GetDatabaseDriver(), config.GetDataSourceName())
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()
	return db
}
