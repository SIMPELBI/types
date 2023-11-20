package model

import "github.com/golang-module/carbon/v2"

type AmiDokumenSpmi struct {
	IDDokumen   int           `gorm:"primaryKey;column:id_dokumen" json:"-"`
	Judul       string        `gorm:"column:judul" json:"judul"`
	Dokumen     string        `gorm:"column:dokumen" json:"dokumen"`
	Tanggal     carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin     int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode   int           `gorm:"column:id_periode" json:"id_periode"`
	LinkDokumen string        `gorm:"column:link_dokumen" json:"link_dokumen"`
}

type AmiDokumenSpmiJoin struct {
	IDDokumen   int           `gorm:"primaryKey;column:id_dokumen" json:"-"`
	Judul       string        `gorm:"column:judul" json:"judul"`
	Dokumen     string        `gorm:"column:dokumen" json:"dokumen"`
	Tanggal     carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin     int           `gorm:"column:id_admin" json:"id_admin"`
	NmAdmin     string        `gorm:"column:nm_admin" json:"nm_admin"`
	IDPeriode   int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun       string        `gorm:"column:tahun" json:"tahun"`
	LinkDokumen string        `gorm:"column:link_dokumen" json:"link_dokumen"`
}

func (m *AmiDokumenSpmi) TableName() string {
	return "ami_dokumen_spmi"
}