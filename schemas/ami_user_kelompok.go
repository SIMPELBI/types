package model

type AmiUserKelompok struct {
	IDKelompok   int    `gorm:"primaryKey;column:id_kelompok" json:"-"`
	NamaKelompok string `gorm:"column:nama_kelompok" json:"nama_kelompok"`
	KodeUnit     string `gorm:"column:kode_unit" json:"kode_unit"`
}

func (m *AmiUserKelompok) TableName() string {
	return "ami_user_kelompok"
}
