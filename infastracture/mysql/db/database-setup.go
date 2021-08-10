package db

import (
	"bvapi/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func CreateTables() {
	fmt.Println("start creating table")
	db, err := repository.Repository()
	fmt.Println("establishing connection ...")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	fmt.Println("connection established..")
	fmt.Println("start creating table...")
	createUserTable(db)
	fmt.Println("table created successfully")
}
func createUserTable(db *sql.DB) {
	fmt.Println("creating query...")
	query := `CREATE TABLE IF NOT EXISTS user(id varchar(25) primary key, contact varchar(25), user_names varchar(25), surname varchar(25),date_of_birth datetime default CURRENT_TIMESTAMP,bio longblob)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	fmt.Println("executing query...")
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
	}
	fmt.Println("query executed")
}
