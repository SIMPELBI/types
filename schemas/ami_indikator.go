package model

type AmiIndikator struct {
	IDIndikator   int    `gorm:"primaryKey;column:id_indikator" json:"-"`
	NamaIndikator string `gorm:"column:nama_indikator" json:"nama_indikator"`
	Isi           string `gorm:"column:isi" json:"isi"`
}

func (m *AmiIndikator) TableName() string {
	return "ami_indikator"
}
