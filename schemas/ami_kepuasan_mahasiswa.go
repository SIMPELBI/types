package model

import "github.com/golang-module/carbon/v2"

type AmiKepuasanMahasiswa struct {
	IDKepuasanMahasiswa int           `gorm:"primaryKey;column:id_kepuasan_mahasiswa" json:"-"`
	Judul               string        `gorm:"column:judul" json:"judul"`
	File                string        `gorm:"column:file" json:"file"`
	Tanggal             carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin             int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode           int           `gorm:"column:id_periode" json:"id_periode"`
}
