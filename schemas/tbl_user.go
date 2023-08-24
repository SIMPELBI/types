package model

import "database/sql"

type TblUser struct {
	IDUsers              int            `gorm:"primaryKey;column:id_users" json:"-"`
	IDUserLevel          int            `gorm:"column:id_user_level" json:"id_user_level"`
	IsAktif              string         `gorm:"column:is_aktif" json:"is_aktif"`
	FullName             sql.NullString `gorm:"column:full_name" json:"full_name"`
	Email                sql.NullString `gorm:"column:email" json:"email"`
	Password             sql.NullString `gorm:"column:password" json:"password"`
	NomorTelepon         sql.NullString `gorm:"column:nomor_telepon" json:"nomor_telepon"`
	Images               sql.NullString `gorm:"column:images" json:"images"`
	IDSiap               sql.NullString `gorm:"column:id_siap" json:"id_siap"`
	GoogleToken          sql.NullString `gorm:"column:google_token" json:"google_token"`
	JabatanID            sql.NullInt32  `gorm:"column:jabatan_id" json:"jabatan_id"`
	KelompokID           sql.NullInt32  `gorm:"column:kelompok_id" json:"kelompok_id"`
	ProdiID              sql.NullString `gorm:"column:prodi_id" json:"prodi_id"`
	PenanggungjawabSurat sql.NullInt32  `gorm:"column:penanggungjawab_surat" json:"penanggungjawab_surat"`
	Statusdosenid        sql.NullString `gorm:"column:StatusDosenID" json:"status_dosen_id"`
}

func (m *TblUser) TableName() string {
	return "tbl_user"
}
