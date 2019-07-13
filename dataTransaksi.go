package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAllTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi Transaksi
	var arr_transaksi []Transaksi
	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	rows, err := db.Query("select * from transaksi")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&transaksi.IdTransaksi,
			&transaksi.Nama_Tes,
			&transaksi.Keterangan,
			&transaksi.Tanggal_Tes,
			&transaksi.Lokasi,
			&transaksi.Peminta_Tes,
			&transaksi.Customer,
			&transaksi.Finance_User,
			&transaksi.Finance_Biaya,
			&transaksi.Finance_Tgl,
			&transaksi.File,
			&transaksi.Status,
			&transaksi.IdUser,
		); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_transaksi = append(arr_transaksi, transaksi)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Result = arr_transaksi
	log.Print("Request Data Transaksi")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
