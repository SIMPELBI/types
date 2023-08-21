package model

import "database/sql"

type AmiAdmin struct {
	IDAdmin  int            `gorm:"primaryKey;column:id_admin" json:"-"`
	UID      string         `gorm:"primaryKey;column:uid" json:"-"`
	NmAdmin  string         `gorm:"column:nm_admin" json:"nm_admin"`
	Password string         `gorm:"column:password" json:"password"`
	Nidn     string         `gorm:"column:nidn" json:"nidn"`
	Jabatan  sql.NullString `gorm:"column:jabatan" json:"jabatan"`
	Email    sql.NullString `gorm:"column:email" json:"email"`
	Foto     sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiAdmin) TableName() string {
	return "ami_admin"
}
