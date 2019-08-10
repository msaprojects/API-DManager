package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root1@tcp(127.0.0.1:3306)/document_manager")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func conn488() *sql.DB {
	db, err := sql.Open("mysql", "root:hanyaadminyangtau@tcp(192.168.4.77:3306)/c_erp_sigk")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
