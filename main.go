package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	http.Handle("/", router)

	// USER
	router.HandleFunc("/getusers", getAllUsers).Methods("GET")
	router.HandleFunc("/user", insertUser).Methods("POST")
	router.HandleFunc("/user", updateUser).Methods("PUT")
	router.HandleFunc("/user", deleteUser).Methods("DELETE")

	// TRANSAKSI
	// router.HandleFunc("/gettransaksi", getAllTransaksi).Methods("GET")
	// router.HandleFunc("/transaksi", insertTransaksi).Methods("POST")
	// router.HandleFunc("/transaksi", updateTransaksi).Methods("PUT")
	// router.HandleFunc("/transaksi", deleteTransaksi).Methods("DELETE")
	// http.Handle("/", router)

	//STARTING LOG
	fmt.Println("Aplikasi Document Manager port :8989")
	log.Fatal(http.ListenAndServe(":8989", router))

}
