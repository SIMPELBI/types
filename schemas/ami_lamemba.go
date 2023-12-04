package model

import "github.com/golang-module/carbon/v2"

type AmiLamemba struct {
	IDLamemba        int           `gorm:"primaryKey;column:id_lamemba" json:"-"`
	IDSkSertifikat   int           `gorm:"column:id_sk_sertifikat" json:"id_sk_sertifikat"`
	Peringkat        string        `gorm:"column:peringkat" json:"peringkat"`
	Konten           string        `gorm:"column:konten" json:"konten"`
	MasaBerlakuStart carbon.Carbon `gorm:"column:masa_berlaku_start" json:"masa_berlaku_start"`
	MasaBerlakuEnd   carbon.Carbon `gorm:"column:masa_berlaku_end" json:"masa_berlaku_end"`
	IDJenjang        int           `gorm:"column:id_jenjang" json:"id_jenjang"`
	IDProdi          int           `gorm:"column:id_prodi" json:"id_prodi"`
	IDAdmin          int           `gorm:"column:id_admin" json:"id_admin"`
}

func (m *AmiLamemba) TableName() string {
	return "ami_ban_pt"
}
