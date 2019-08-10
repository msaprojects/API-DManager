package main

import "database/sql"

type Dashboard struct {
	JMLBERKAS       int `from:"jmlberkas" json:"jmlberkas"`
	JMLPUBLISH      int `from:"jmlpublish" json:"jmlpublish"`
	JMLTUGASFINANCE int `from:"jmltgsfinance" json:"jmltgsfinance"`
}

type Customers struct {
	IDCUSTOMER int    `from:"idcustomer" json:"idcustomer"`
	NAMA       string `from:"nama" json:"nama"`
	ALAMAT     string `from:"alamat" json:"alamat"`
	NOTELP     string `from:"notelp" json:"notelp"`
	CP         string `from:"cp" json:"cp"`
	KODESISTEM string `from:"kodesistem" json:"kodesistem"`
	AKTIF      int    `from:"aktif" json:"aktif"`
}

type Users struct {
	IdUser   int    `form:"iduser" json:"iduser"`
	Nama     string `form:"nama" json:"nama"`
	Jabatan  string `form:"jabatan" json:"jabatan"`
	Password string `form:"password" json:"password"`
	Aktif    int    `form:"aktif" json:"aktif"`
}

type Transaksi struct {
	IdTransaksi   string         `form:"idtransaksi" json:"idtransaksi"`
	Nama_Tes      string         `form:"nama_tes" json:"nama_tes"`
	Keterangan    string         `form:"keterangan" json:"keterangan"`
	Tanggal_Tes   string         `form:"tgl_tes" json:"tgl_tes"`
	Lokasi        string         `form:"lokasi" json:"lokasi"`
	Peminta_Tes   string         `form:"peminta_tes" json:"peminta_tes"`
	Customer      string         `form:"customer" json:"customer"`
	Finance_User  string         `form:"finance_user" json:"finance_user"`
	Finance_Tgl   string         `form:"finance_tgl" json:"finance_tgl"`
	Finance_Biaya string         `form:"finance_biaya" json:"finance_biaya"`
	File          sql.NullString `form:"file" json:"file"`
	Status        string         `form:"status" json:"status"`
	IdUser        int            `form:"iduser" json:"iduser"`
	IDCUSTOMER    int            `form:"idcustomer" json:"idcustomer"`
}

type ResponseUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []Users
}

type ResponseTransaksi struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []Transaksi
}

type ResponseCustomer struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []Customers
}

type ResponseDashboard struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []Dashboard
}
