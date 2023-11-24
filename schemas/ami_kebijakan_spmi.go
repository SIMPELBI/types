package model

import "github.com/golang-module/carbon/v2"

type AmiKebijakanSpmi struct {
	IDKebijakanSpmi int           `gorm:"primaryKey;column:id_kebijakan_spmi" json:"-"`
	Judul           string        `gorm:"column:judul" json:"judul"`
	Keterangan      string        `gorm:"column:keterangan" json:"keterangan"`
	File            string        `gorm:"column:file" json:"file"`
	LinkDokumen     string        `gorm:"column:link_dokumen" json:"link_dokumen"`
	Tanggal         carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin         int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode       int           `gorm:"column:id_periode" json:"id_periode"`
}

type AmiKebijakanSpmiJoin struct {
	IDKebijakanSpmi int           `gorm:"primaryKey;column:id_kebijakan_spmi" json:"-"`
	Judul           string        `gorm:"column:judul" json:"judul"`
	Keterangan      string        `gorm:"column:keterangan" json:"keterangan"`
	File            string        `gorm:"column:file" json:"file"`
	LinkDokumen     string        `gorm:"column:link_dokumen" json:"link_dokumen"`
	Tanggal         carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin         int           `gorm:"column:id_admin" json:"id_admin"`
	NamaAdmin       string        `gorm:"column:nm_admin" json:"nm_admin"`
	IDPeriode       int           `gorm:"column:id_periode" json:"id_periode"`
	Tahun           string        `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiKebijakanSpmi) TableName() string {
	return "ami_kebijakan_spmi"
}
