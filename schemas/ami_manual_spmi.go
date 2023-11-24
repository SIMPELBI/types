package model

import "github.com/golang-module/carbon/v2"

type AmiManualSpmi struct {
	IDManualSpmi int           `gorm:"primaryKey;column:id_manual_spmi" json:"-"`
	Judul        string        `gorm:"column:judul" json:"judul"`
	Keterangan   string        `gorm:"column:keterangan" json:"keterangan"`
	File         string        `gorm:"column:file" json:"file"`
	Tanggal      carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin      int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode    int           `gorm:"column:id_periode" json:"id_periode"`
}

type AmiManualSpmiJoin struct {
	IDManualSpmi int           `gorm:"primaryKey;column:id_manual_spmi" json:"-"`
	Judul        string        `gorm:"column:judul" json:"judul"`
	Keterangan   string        `gorm:"column:keterangan" json:"keterangan"`
	File         string        `gorm:"column:file" json:"file"`
	Tanggal      carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin      int           `gorm:"column:id_admin" json:"id_admin"`
	NamaAdmin    string        `gorm:"column:nm_admin" json:"nm_admin"`
	IDPeriode    int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun        string        `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiManualSpmi) TableName() string {
	return "ami_manual_spmi"
}
