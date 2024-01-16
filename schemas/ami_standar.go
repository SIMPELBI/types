package model

type AmiStandar struct {
	IDStandar   int    `gorm:"primaryKey;column:id_standar" json:"-"`
	IDIndikator int    `gorm:"column:id_indikator" json:"id_indikator"`
	Standar     string `gorm:"column:standar" json:"standar"`
	Isi         string `gorm:"column:isi" json:"isi"`
	IDProdiUnit int    `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
	IDSiklus    int    `gorm:"column:id_siklus" json:"id_siklus"`
}

type AmiStandarJoin struct {
	IDStandar     int    `gorm:"primaryKey;column:id_standar" json:"-"`
	Standar       string `gorm:"column:standar" json:"standar"`
	Isi           string `gorm:"column:isi" json:"isi"`
	IDIndikator   int    `gorm:"column:id_indikator" json:"id_indikator"`
	NamaIndikator string `gorm:"column:nama_indikator" json:"nama_indikator"`
	IsiIndikator  string `gorm:"column:isi" json:"isi_indikator"`
	IDProdiUnit   int    `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
	ProdiUnit     string `gorm:"column:prodi_unit" json:"prodi_unit"`
	IDSiklus      int    `gorm:"column:id_siklus" json:"id_siklus"`
	Tahun         int    `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiStandar) TableName() string {
	return "ami_standar"
}
