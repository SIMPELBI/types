package model

import "database/sql"

type AmiHakAkses struct {
	IDHakAkses  int           `gorm:"primaryKey;column:id_hak_akses" json:"-"`
	IDUserLevel sql.NullInt32 `gorm:"column:id_user_level" json:"id_user_level"`
	IDMenu      sql.NullInt32 `gorm:"column:id_menu" json:"id_menu"`
}

func (m *AmiHakAkses) TableName() string {
	return "ami_hak_akses"
}
