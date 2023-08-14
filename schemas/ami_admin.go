package model

import "database/sql"

type AmiAdmin struct {
	IDAdmin  int            `gorm:"primaryKey;column:id_admin" json:"-"`
	NmAdmin  string         `gorm:"column:nm_admin" json:"nmAdmin"`
	Jabatan  sql.NullString `gorm:"column:jabatan" json:"jabatan"`
	Email    sql.NullString `gorm:"column:email" json:"email"`
	Nidn     string         `gorm:"column:nidn" json:"nidn"`
	Password string         `gorm:"column:password" json:"password"`
	Foto     sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus sql.NullInt32  `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiAdmin) TableName() string {
	return "ami_admin"
}
