package model

import "github.com/golang-module/carbon/v2"

type AmiSkSertifikat struct {
	IDSkSertifikat   int           `gorm:"primaryKey;column:id_sk_sertifikat" json:"-"`
	NamaSkSertifikat string        `gorm:"column:nama_sk_sertifikat" json:"nama_sk_sertifikat"`
	File             string        `gorm:"column:file" json:"file"`
	Tanggal          carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin          int           `gorm:"column:id_admin" json:"id_admin"`
}
