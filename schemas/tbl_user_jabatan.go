package model

type AmiUserJabatan struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	Singkatan string `gorm:"column:Singkatan" json:"Singkatan"`
	Nama      string `gorm:"column:Nama" json:"Nama"`
	KodeUnit  string `gorm:"column:kode_unit" json:"kode_unit"`
}

func (m *AmiUserJabatan) TableName() string {
	return "tbl_user_jabatan"
}
