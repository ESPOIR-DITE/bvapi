package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Repository() (*sql.DB, error) {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bvapi")
	//defer db.Close()
	if err != nil {

		fmt.Println("error conecting to the database")
		log.Fatal(err)
		return db, err
	}
	return db, nil
}
