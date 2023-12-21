package model

type AmiMasterUsers struct {
	IDUsers      int    `gorm:"primaryKey;column:id_users" json:"-"`
	FullName     string `gorm:"column:full_name" json:"full_name"`
	Email        string `gorm:"column:email" json:"email"`
	Password     string `gorm:"column:password" json:"password"`
	NomorTelepon string `gorm:"column:nomor_telepon" json:"nomor_telepon"`
	Foto         string `gorm:"column:foto" json:"foto"`
	IDUserLevel  int    `gorm:"column:id_user_level" json:"id_user_level"`
	IDSiap       string `gorm:"column:id_siap" json:"id_siap"`
	IDJabatan    int    `gorm:"column:id_jabatan" json:"id_jabatan"`
	IDKelompok   int    `gorm:"column:id_kelompok" json:"id_kelompok"`
}

func (m *AmiMasterUsers) TableName() string {
	return "ami_master_users"
}
