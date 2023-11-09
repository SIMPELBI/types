package model

import "github.com/golang-module/carbon/v2"

type AmiSpmiTentang struct {
	IDTentang int           `gorm:"primaryKey;column:id_tentang" json:"-"`
	Judul     string        `gorm:"column:judul" json:"judul"`
	Isi       string        `gorm:"column:isi" json:"isi"`
	Gambar    string        `gorm:"column:gambar" json:"gambar"`
	Tanggal   carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin   int           `gorm:"column:id_admin" json:"id_admin"`
	NmAdmin   string        `gorm:"column:nm_admin" json:"nm_admin,omitempty"`
	IDPeriode int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun     string        `gorm:"column:tahun" json:"tahun,omitempty"`
}

func (m *AmiSpmiTentang) TableName() string {
	return "ami_spmi_tentang"
}
