package model

type TblUserKelompok struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`
	NamaKelompok string `gorm:"column:nama_kelompok" json:"nama_kelompok"`
	KodeUnit     string `gorm:"column:kode_unit" json:"kode_unit"`
}

func (m *TblUserKelompok) TableName() string {
	return "tbl_user_kelompok"
}
