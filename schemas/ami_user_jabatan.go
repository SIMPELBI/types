package model

type AmiUserJabatan struct {
	IDJabatan int    `gorm:"primaryKey;column:id_jabatan" json:"-"`
	Singkatan string `gorm:"column:singkatan" json:"singkatan"`
	Nama      string `gorm:"column:nama" json:"nama"`
	KodeUnit  string `gorm:"column:kode_unit" json:"kode_unit"`
}

func (m *AmiUserJabatan) TableName() string {
	return "ami_user_jabatan"
}
