package model

type AmiStandar struct {
	IDStandar   int    `gorm:"primaryKey;column:id_standar" json:"-"`
	IDIndikator int    `gorm:"column:id_indikator" json:"id_indikator"`
	Standar     string `gorm:"column:standar" json:"standar"`
	Isi         string `gorm:"column:isi" json:"isi"`
	IDProdiUnit int    `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
}

type AmiStandarJoin struct {
	IDStandar     int    `gorm:"primaryKey;column:id_standar" json:"-"`
	IDIndikator   int    `gorm:"column:id_indikator" json:"id_indikator"`
	NamaIndikator string `gorm:"column:nama_indikator" json:"nama_indikator"`
	Standar       string `gorm:"column:standar" json:"standar"`
	Isi           string `gorm:"column:isi" json:"isi"`
	IDProdiUnit   int    `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
	ProdiUnit     string `gorm:"column:prodi_unit" json:"prodi_unit"`
}

func (m *AmiStandar) TableName() string {
	return "ami_standar"
}
