package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Aplikasi Document Manager port :8989")
	http.HandleFunc("/login", Login)
	http.HandleFunc("/dashboard", getDashboard)
	http.HandleFunc("/transaksidashboard", getAllTransaksiDashboard)
	http.HandleFunc("/transaksiblmaccfinance", getAllTransaksiBlmAccFinance)
	// DATA
	http.HandleFunc("/user", getAllUsers)
	http.HandleFunc("/useri", UserInsert)
	http.HandleFunc("/useru", UserUpdate)
	http.HandleFunc("/userd", UserDelete)
	// DATA TRANSAKSI
	http.HandleFunc("/transaksi", getAllTransaksi)
	http.HandleFunc("/transaksii", TambahTransaksi)
	http.HandleFunc("/transaksiu", UbahTransaksi)
	http.HandleFunc("/transaksid", DeleteTransaksi)
	// DATA CUSTOMER
	http.HandleFunc("/customer", getAllCustomer)
	http.HandleFunc("/customeri", CustomerInsert)
	http.HandleFunc("/customeru", CustomerUpdate)
	http.HandleFunc("/customerd", CustomerDelete)

	http.ListenAndServe(":8989", nil)
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	var d Dashboard
	var arr_d []Dashboard
	var response ResponseDashboard

	db := connect()
	defer db.Close()

	rows, err := db.Query("select count(idtransaksi) as jmlberkas, (select count(idtransaksi) from transaksi where status='Publish') as jmlpublish, (select count(idtransaksi) from transaksi where finance_biaya='' || finance_biaya='0' ) as jmltgsfinance from transaksi;")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&d.JMLBERKAS,
			&d.JMLPUBLISH,
			&d.JMLTUGASFINANCE,
		); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_d = append(arr_d, d)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Result = arr_d
	log.Print("Request Data Dashboard")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getAllTransaksiDashboard(w http.ResponseWriter, r *http.Request) {
	var transaksi Transaksi
	var arr_transaksi []Transaksi
	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	rows, err := db.Query("select f.idtransaksi, f.nama_tes, f.keterangan, f.tgl_tes, f.lokasi, f.peminta_tes, c.nama as customer, f.finance_user, f.finance_biaya, f.finance_tgl, f.file, f.status, f.iduser from transaksi f, customer c where f.idcustomer=c.idcustomer and status = 'Publish';")
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

func getAllTransaksiBlmAccFinance(w http.ResponseWriter, r *http.Request) {
	var transaksi Transaksi
	var arr_transaksi []Transaksi
	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	rows, err := db.Query("select f.idtransaksi, f.nama_tes, f.keterangan, f.tgl_tes, f.lokasi, f.peminta_tes, c.nama as customer, f.finance_user, f.finance_biaya, f.finance_tgl, f.file, f.status, f.iduser from transaksi f, customer c where f.idcustomer=c.idcustomer and finance_biaya='';")
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

func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	var user Users
	var arr_user []Users
	// var response ResponseUser

	err := decoder.Decode(&user)
	log.Println("adudu ", user.Nama, user.Password)
	rows, err := db.Query("Select iduser, nama, jabatan From user where nama = ? and password = ? and aktif = 1",
		user.Nama,
		user.Password,
	)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(
			&user.IdUser,
			&user.Nama,
			&user.Jabatan,
		); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, user)
		}
	}

	// response.Status = 1
	// response.Message = "Success"
	// response.Result = arr_user
	log.Print("User Login")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arr_user)
}

// DATA USER
// GET DATA USER
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var user Users
	var arr_user []Users
	var response ResponseUser

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select iduser, nama, jabatan, password, aktif From user order by iduser asc")
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

// INSERT DATA USER
func UserInsert(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var u Users
	var response ResponseUser
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO user (nama, jabatan, password, aktif) VALUES (?, ?, ?, ?)",
		u.Nama,
		u.Jabatan,
		u.Password,
		u.Aktif,
	)
	response.Status = 1
	response.Message = "Success Add"
	log.Print("Tambah Data User")

}

// UPDATE DATA USER
func UserUpdate(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var u Users
	var response ResponseUser
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("UPDATE user set nama = ?, jabatan = ?, password = ?, aktif = ? where iduser = ?",
		u.Nama,
		u.Jabatan,
		u.Password,
		u.Aktif,
		u.IdUser,
	)
	response.Status = 1
	response.Message = "Success Add"
	log.Print("Update Data User")
}

// DELETE DATA USER
func UserDelete(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var u Users
	var response ResponseUser
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM user where iduser = ?",
		u.IdUser,
	)
	response.Status = 1
	response.Message = "Success Delete"
	log.Print("Delete Data User")
}

// END DATA USER

//START TRANSAKSI
// GET ALL DATA TRANSAKSI
func getAllTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi Transaksi
	var arr_transaksi []Transaksi
	var response ResponseTransaksi

	db := connect()
	defer db.Close()

	rows, err := db.Query("select t.idtransaksi, t.nama_tes, t.keterangan, t.tgl_tes, t.lokasi, t.peminta_tes, t.finance_user, t.finance_biaya, t.finance_tgl, t.file, t.status, t.iduser, t.idcustomer, c.nama from transaksi t, customer c where t.idcustomer=c.idcustomer;")
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
			&transaksi.Finance_User,
			&transaksi.Finance_Biaya,
			&transaksi.Finance_Tgl,
			&transaksi.File,
			&transaksi.Status,
			&transaksi.IdUser,
			&transaksi.IDCUSTOMER,
			&transaksi.Customer,
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

//TAMBAH DATA TRANSAKSI
func TambahTransaksi(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var t Transaksi
	var response ResponseTransaksi
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO transaksi (nama_tes, keterangan, tgl_tes, lokasi, peminta_tes, finance_user, finance_tgl, finance_biaya, status, iduser, idcustomer) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		t.Nama_Tes,
		t.Keterangan,
		t.Tanggal_Tes,
		t.Lokasi,
		t.Peminta_Tes,
		t.Finance_User,
		t.Finance_Tgl,
		t.Finance_Biaya,
		t.Status,
		t.IdUser,
		t.IDCUSTOMER,
	)

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Tambah Data Transaksi", err)
}

// DELETE DATA TRANSAKSI
func DeleteTransaksi(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var t Transaksi
	var response ResponseTransaksi
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM transaksi where idtransaksi = ?",
		t.IdTransaksi,
	)
	response.Status = 1
	response.Message = "Success Delete"
	log.Print("Delete Data Transaksi")
}

// UBAH DATA TRANSAKSI
func UbahTransaksi(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var t Transaksi
	var response ResponseTransaksi
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("UPDATE transaksi setnama_tes = ?, keterangan = ?, tgl_tes = ?, lokasi = ?, peminta_tes = ?, customer = ?, finance_user = ?, finance_tgl = ?, finance_biaya = ?, file = ?, status = ? WHERE iduser = ?",
		t.Nama_Tes,
		t.Keterangan,
		t.Tanggal_Tes,
		t.Lokasi,
		t.Peminta_Tes,
		t.Customer,
		t.Finance_User,
		t.Finance_Tgl,
		t.Finance_Biaya,
		t.File,
		t.Status,
		t.IdUser,
	)
	response.Status = 1
	response.Message = "Success Add"
	log.Print("Ubah Data Transaksi")
}

//END TRANSAKSI

// DATA CUSTOMER
// LIST DATA CUSTOMER
func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	var c Customers
	var arr_cus []Customers
	var response ResponseCustomer

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select idcustomer, nama, alamat, notelp, cp, kodesistem, aktif From customer where aktif=1")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(
			&c.IDCUSTOMER,
			&c.NAMA,
			&c.ALAMAT,
			&c.NOTELP,
			&c.CP,
			&c.KODESISTEM,
			&c.AKTIF,
		); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_cus = append(arr_cus, c)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Result = arr_cus
	log.Print("Request Data Customer")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// INSERT DATA CUSTOMER
func CustomerInsert(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var c Customers
	var response ResponseCustomer
	err := decoder.Decode(&c)

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO customer (nama, alamat, notelp, cp, kodesistem, aktif) VALUES (?, ?, ?, ?, ?, ?)",
		c.NAMA,
		c.ALAMAT,
		c.NOTELP,
		c.CP,
		c.KODESISTEM,
		c.AKTIF,
	)
	response.Status = 1
	response.Message = "Success Add"
	log.Print("Tambah Data Customer")
}

// UPDATE DATA CUSTOMER
func CustomerUpdate(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var c Customers
	var response ResponseCustomer
	err := decoder.Decode(&c)

	if err != nil {
		panic(err)
	}
	_, err = db.Exec("UPDATE customer set nama = ?, alamat = ?, notelp = ?, cp = ?, kodesistem = ?, aktif = ? WHERE idcustomer = ?",
		c.NAMA,
		c.ALAMAT,
		c.NOTELP,
		c.CP,
		c.KODESISTEM,
		c.AKTIF,
		c.IDCUSTOMER,
	)
	response.Status = 1
	response.Message = "Success Add"
	log.Print("Update Data Customer", err)
}

// DELETE DATA CUSTOMER
func CustomerDelete(w http.ResponseWriter, request *http.Request) {
	db := connect()
	defer db.Close()

	decoder := json.NewDecoder(request.Body)
	var c Customers
	var response ResponseCustomer
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM customer where idcustomer = ?",
		c.IDCUSTOMER,
	)
	response.Status = 1
	response.Message = "Success Delete"
	log.Print("Delete Data CUstomer")
}
