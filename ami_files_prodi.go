package types

import "github.com/golang-module/carbon/v2"

type AmiFilesProdi struct {
	IDFile   int           `gorm:"primaryKey;column:id_file" json:"-"`
	IDSiklus int           `gorm:"column:id_siklus" json:"idSiklus"`
	Judul    string        `gorm:"column:judul" json:"judul"`
	File     string        `gorm:"column:file" json:"file"`
	Tgl      carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAdmin  int           `gorm:"column:id_admin" json:"idAdmin"`
}

func (m *AmiFilesProdi) TableName() string {
	return "ami_files_prodi"
}
