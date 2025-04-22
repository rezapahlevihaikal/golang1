package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_connect"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal konek ke MySQL: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ping ke DB gagal: %v", err)
	}

	fmt.Println("✅ Database terkoneksi!")
}
