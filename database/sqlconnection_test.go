package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func sqltest(t *testing.T, err error) {
	CFG := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "GoRestApi",
	}

	db, err = sql.Open("mysql", CFG.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	ping := db.Ping()

	if ping != nil {
		log.Fatal(ping)
	}

	fmt.Println("Connected successfully!!")

}
