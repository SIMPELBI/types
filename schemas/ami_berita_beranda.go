package model

import "github.com/golang-module/carbon/v2"

type AmiBeritaSpmi struct {
	IDBeritaBeranda int           `gorm:"primaryKey;column:id_berita_beranda" json:"-"`
	Judul           string        `gorm:"column:judul" json:"judul"`
	Isi             string        `gorm:"column:isi" json:"isi"`
	Gambar          string        `gorm:"column:gambar" json:"gambar"`
	Tanggal         carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin         int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode       int           `gorm:"column:id_periode" json:"id_periode"`
}

func (m *AmiBeritaSpmi) TableName() string {
	return "ami_berita_spmi"
}
