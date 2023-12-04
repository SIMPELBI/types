package model

import "github.com/golang-module/carbon/v2"

type AmiDokumenSpmi struct {
	IDDokumen   int           `gorm:"primaryKey;column:id_dokumen" json:"-"`
	NamaDokumen string        `gorm:"column:nama_dokumen" json:"nama_dokumen"`
	File        string        `gorm:"column:file" json:"file"`
	Tanggal     carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin     int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode   int           `gorm:"column:id_periode" json:"id_periode"`
	Keterangan  string        `gorm:"column:keterangan" json:"keterangan"`
}

type AmiDokumenSpmiJoin struct {
	IDDokumen   int           `gorm:"primaryKey;column:id_dokumen" json:"-"`
	NamaDokumen string        `gorm:"column:nama_dokumen" json:"nama_dokumen"`
	File        string        `gorm:"column:file" json:"file"`
	Tanggal     carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin     int           `gorm:"column:id_admin" json:"id_admin"`
	NmAdmin     string        `gorm:"column:nm_admin" json:"nm_admin"`
	IDPeriode   int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun       string        `gorm:"column:tahun" json:"tahun"`
	Keterangan  string        `gorm:"column:keterangan" json:"keterangan"`
}

func (m *AmiDokumenSpmi) TableName() string {
	return "ami_dokumen_spmi"
}
