package model

import "github.com/golang-module/carbon/v2"

type AmiKebijakanSpmi struct {
	IDKebijakanSpmi int           `gorm:"primaryKey;column:id_kebijakan_spmi" json:"-"`
	Subject         string        `gorm:"column:subject" json:"subject"`
	Keterangan      string        `gorm:"column:keterangan" json:"keterangan"`
	File            string        `gorm:"column:file" json:"file"`
	LinkDokumen     string        `gorm:"column:link_dokumen" json:"link_dokumen"`
	Tanggal         carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin         int           `gorm:"column:id_admin" json:"id_admin"`
	IDPeriode       int           `gorm:"column:id_periode" json:"id_periode"`
}
