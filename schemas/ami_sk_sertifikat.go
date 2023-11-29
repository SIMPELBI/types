package model

import "github.com/golang-module/carbon/v2"

type AmiSkSertifikat struct {
	IDSkSertifikat   int           `gorm:"primaryKey;column:id_sk_sertifikat" json:"-"`
	NamaSkSertifikat string        `gorm:"column:nama_sk_sertifikat" json:"nama_sk_sertifikat"`
	File             string        `gorm:"column:file" json:"file"`
	Tanggal          carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin          int           `gorm:"column:id_admin" json:"id_admin"`
}

type AmiSkSertifikatJoin struct {
	IDSkSertifikat   int           `gorm:"primaryKey;column:id_sk_sertifikat" json:"-"`
	NamaSkSertifikat string        `gorm:"column:nama_sk_sertifikat" json:"nama_sk_sertifikat"`
	File             string        `gorm:"column:file" json:"file"`
	Tanggal          carbon.Carbon `gorm:"column:tanggal" json:"tanggal"`
	IDAdmin          int           `gorm:"column:id_admin" json:"id_admin"`
	NamaAdmin        string        `gorm:"column:nm_admin" json:"nm_admin"`
}

func (m *AmiSkSertifikat) TableName() string {
	return "ami_sk_sertifikat"
}
