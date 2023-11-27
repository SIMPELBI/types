package model

import "github.com/golang-module/carbon/v2"

type AmiStandarSpmi struct {
	IDStandarSpmi int           `gorm:"primaryKey;column:id_standar_spmi" json:"-"`
	Judul         string        `gorm:"column:judul" json:"judul"`
	Keterangan    string        `gorm:"column:keterangan" json:"keterangan"`
	File          string        `gorm:"column:file" json:"file"`
	Tanggal       carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin       int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode     int           `gorm:"column:id_periode" json:"id_periode"`
}

func (m *AmiStandarSpmi) TableName() string {
	return "ami_standar_spmi"
}
