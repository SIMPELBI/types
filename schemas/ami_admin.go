package model

import "database/sql"

type AmiAdmin struct {
	IDAdmin    int            `gorm:"primaryKey;column:id_admin" json:"-"`
	UserName   string         `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert     `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	NmAdmin    string         `gorm:"column:nm_admin" json:"nm_admin"`
	Jabatan    sql.NullString `gorm:"column:jabatan" json:"jabatan"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
	IDLevel    int            `gorm:"column:id_level" json:"id_level"`
}

func (m *AmiAdmin) TableName() string {
	return "ami_admin"
}
