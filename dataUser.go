package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var user Users
	var arr_user []Users
	var response ResponseUser

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select iduser, nama, jabatan, password, aktif From user")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&user.IdUser,
			&user.Nama, &user.Jabatan,
			&user.Password,
			&user.Aktif,
		); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, user)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Result = arr_user
	log.Print("Request Data User")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUser(w http.ResponseWriter, r *http.Request) {

	var response ResponseUser

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(2048)
	if err != nil {
		panic(err)
	}

	nama := r.FormValue("nama")
	jabatan := r.FormValue("jabatan")
	password := r.FormValue("password")
	aktif := r.FormValue("aktif")

	_, err = db.Exec("INSERT INTO user (nama, jabatan, password, aktif) VALUES (?, ?, ?, ?)",
		nama,
		jabatan,
		password,
		aktif,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Tambah data user ke database Berhasil")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	var response ResponseUser

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(2048)
	if err != nil {
		panic(err)
	}

	iduser := r.FormValue("iduser")
	nama := r.FormValue("nama")
	jabatan := r.FormValue("jabatan")
	password := r.FormValue("password")
	aktif := r.FormValue("aktif")

	_, err = db.Exec("UPDATE user set nama = ?, jabatan = ?, password = ?, aktif = ? where iduser = ?",
		nama,
		jabatan,
		password,
		aktif,
		iduser,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"
	log.Print("Update data user ke database Berhasil")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	var response ResponseUser

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(2048)
	if err != nil {
		panic(err)
	}

	iduser := r.FormValue("iduser")
	_, err = db.Exec("Delete From user where iduser = ?",
		iduser,
	)

	if err != nil {
		log.Print(err)
		log.Print("Delete data user dari databaser Gagal")
	}

	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data user dari database Berhasil")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
