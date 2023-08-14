package model

import "database/sql"

type AmiStandar struct {
	IDStandar  int            `gorm:"primaryKey;column:id_standar" json:"-"`
	Standar    sql.NullString `gorm:"column:standar" json:"standar"`
	UtkPilihan string         `gorm:"column:utk_pilihan" json:"utkPilihan"`
	Isi        string         `gorm:"column:isi" json:"isi"`
	IDSiklus   int            `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiStandar) TableName() string {
	return "ami_standar"
}
