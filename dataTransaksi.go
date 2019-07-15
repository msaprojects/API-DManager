package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// GET ALL DATA TRANSAKSI
func getAllTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi Transaksi
	var arr_transaksi []Transaksi
	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	rows, err := db.Query("select idtransaksi, nama_tes, keterangan, tgl_tes, lokasi, peminta_tes, customer, finance_user, finance_biaya, finance_tgl, file, status, iduser from transaksi")
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

// INSERT DATA TRANSAKSI
func insertTransaksi(w http.ResponseWriter, r *http.Request) {

	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(2048)
	if err != nil {
		panic(err)
	}

	Nama := r.FormValue("nama_tes")
	Keterangan := r.FormValue("keterangan")
	TanggalTes := r.FormValue("tgl_tes")
	Lokasi := r.FormValue("lokasi")
	PemintaTes := r.FormValue("peminta_tes")
	Customer := r.FormValue("customer")
	FinanceUser := r.FormValue("finance_user")
	FinaneBiaya := r.FormValue("finance_biaya")
	FinanceTgl := r.FormValue("finance_tgl")
	File := r.FormValue("file")
	Status := r.FormValue("status")
	IdUser := r.FormValue("iduser")

	_, err = db.Exec("INSERT INTO transaksi (nama_tes, keterangan, tgl_tes, lokasi, peminta_tes, customer, finance_user, finance_tgl, finance_biaya, file, status, iduser) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		Nama,
		Keterangan,
		TanggalTes,
		Lokasi,
		Lokasi,
		PemintaTes,
		Customer,
		FinanceUser,
		FinanceTgl,
		FinaneBiaya,
		File,
		Status,
		IdUser,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Tambah transaksi user ke database Berhasil")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UPDATE DATA TRANSAKSI
func updateTransaksi(w http.ResponseWriter, r *http.Request) {

	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(2048)
	if err != nil {
		panic(err)
	}

	IdTransaksi := r.FormValue("idtransaksi")
	Nama := r.FormValue("nama_tes")
	Keterangan := r.FormValue("keterangan")
	TanggalTes := r.FormValue("tgl_tes")
	Lokasi := r.FormValue("lokasi")
	PemintaTes := r.FormValue("peminta_tes")
	Customer := r.FormValue("customer")
	FinanceUser := r.FormValue("finance_user")
	FinaneBiaya := r.FormValue("finance_biaya")
	FinanceTgl := r.FormValue("finance_tgl")
	File := r.FormValue("file")
	Status := r.FormValue("status")
	IdUser := r.FormValue("iduser")

	_, err = db.Exec("UPDATE transaksi SET nama_tes = ?, keterangan = ?, tgl_tes = ?,  lokasi = ?, peminta_tes = ?, customer = ?, finance_user = ?, finance_biaya = ? finance_tgl = ?, file = ?, status = ?, iduser = ? WHERE idtransaksi = ?",
		Nama,
		Keterangan,
		TanggalTes,
		Lokasi,
		Lokasi,
		PemintaTes,
		Customer,
		FinanceUser,
		FinaneBiaya,
		FinanceTgl,
		File,
		Status,
		IdUser,
		IdTransaksi,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"
	log.Print("Update transaksi user ke database Berhasil")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DELETE DATA USER
func deleteTransaksi(w http.ResponseWriter, r *http.Request) {

	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(2048)
	if err != nil {
		panic(err)
	}

	IdTransaksi := r.FormValue("idtransaksi")
	_, err = db.Exec("Delete From transaksi where idtransaksi = ?",
		IdTransaksi,
	)

	if err != nil {
		log.Print(err)
		log.Print("Delete data transaksi dari databaser Gagal")
	}

	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data transaksi dari database Berhasil")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
