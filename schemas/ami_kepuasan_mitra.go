package model

import "github.com/golang-module/carbon/v2"

type AmiKepuasanMitra struct {
	IDKepuasanMitra int           `gorm:"primaryKey;column:id_kepuasan_mitra" json:"-"`
	Judul           string        `gorm:"column:judul" json:"judul"`
	File            string        `gorm:"column:file" json:"file"`
	Keterangan      string        `gorm:"column:keterangan" json:"keterangan"`
	Tanggal         carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin         int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode       int           `gorm:"column:id_periode" json:"id_periode"`
}

type AmiKepuasanMitraJoin struct {
	IDKepuasanMitra int           `gorm:"primaryKey;column:id_kepuasan_mitra" json:"-"`
	Judul           string        `gorm:"column:judul" json:"judul"`
	File            string        `gorm:"column:file" json:"file"`
	Keterangan      string        `gorm:"column:keterangan" json:"keterangan"`
	Tanggal         carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin         int           `gorm:"column:id_admin" json:"id_admin"`
	NamaAdmin       string        `gorm:"column:nm_admin" json:"nm_admin"`
	IDPeriode       int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun           string        `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiKepuasanMitra) TableName() string {
	return "ami_kepuasan_mitra"
}
