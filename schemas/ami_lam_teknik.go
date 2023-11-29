package model

import "github.com/golang-module/carbon/v2"

type AmiLamTeknik struct {
	IDLamTeknik      int           `gorm:"primaryKey;column:id_lamteknik" json:"-"`
	SkSertifikat     string        `gorm:"column:sk_sertifikat" json:"sk_sertifikat"`
	Peringkat        string        `gorm:"column:peringkat" json:"peringkat"`
	Konten           string        `gorm:"column:konten" json:"konten"`
	MasaBerlakuStart carbon.Carbon `gorm:"column:masa_berlaku_start" json:"masa_berlaku_start"`
	MasaBerlakuEnd   carbon.Carbon `gorm:"column:masa_berlaku_end" json:"masa_berlaku_end"`
	IDJenjang        int           `gorm:"column:id_jenjang" json:"id_jenjang"`
	IDProdi          int           `gorm:"column:id_prodi" json:"id_prodi"`
	IDAdmin          int           `gorm:"column:id_admin" json:"id_admin"`
}

type AmiLamTeknikJoin struct {
	IDLamTeknik      int           `gorm:"primaryKey;column:id_lamteknik" json:"-"`
	SkSertifikat     string        `gorm:"column:sk_sertifikat" json:"sk_sertifikat"`
	Peringkat        string        `gorm:"column:peringkat" json:"peringkat"`
	Konten           string        `gorm:"column:konten" json:"konten"`
	MasaBerlakuStart carbon.Carbon `gorm:"column:masa_berlaku_start" json:"masa_berlaku_start"`
	MasaBerlakuEnd   carbon.Carbon `gorm:"column:masa_berlaku_end" json:"masa_berlaku_end"`
	IDJenjang        int           `gorm:"column:id_jenjang" json:"id_jenjang"`
	Jenjang          string        `gorm:"column:jenjang" json:"jenjang"`
	IDProdi          int           `gorm:"column:id_prodi" json:"id_prodi"`
	Prodi            string        `gorm:"column:prodi" json:"prodi"`
	IDAdmin          int           `gorm:"column:id_admin" json:"id_admin"`
	NamaAdmin        string        `gorm:"column:nm_admin" json:"nm_admin"`
}

func (m *AmiLamTeknik) TableName() string {
	return "ami_lam_teknik"
}
