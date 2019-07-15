package main

import "database/sql"

type Users struct {
	IdUser   string `form:"iduser" json:"iduser"`
	Nama     string `form:"nama" json:"nama"`
	Jabatan  string `form:"jabatan" json:"jabatan"`
	Password string `form:"password" json:"password"`
	Aktif    string `form:"aktif" json:"aktif"`
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
	IdUser        string         `form:"iduser" json:"iduser"`
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
