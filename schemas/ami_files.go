package model

import "github.com/golang-module/carbon/v2"

type AmiFiles struct {
	IDFile   int           `gorm:"primaryKey;column:id_file" json:"-"`
	IDSiklus int           `gorm:"column:id_siklus" json:"id_siklus"`
	Judul    string        `gorm:"column:judul" json:"judul"`
	File     string        `gorm:"column:file" json:"file"`
	Tgl      carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAdmin  int           `gorm:"column:id_admin" json:"id_admin"`
}

func (m *AmiFiles) TableName() string {
	return "ami_files"
}
