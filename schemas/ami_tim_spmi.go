package model

type AmiTimSpmi struct {
	IDAnggota int    `gorm:"primaryKey;column:id_anggota" json:"-"`
	Nama      string `gorm:"column:nama" json:"nama"`
	Jabatan   string `gorm:"column:jabatan" json:"jabatan"`
	Foto      string `gorm:"column:foto" json:"foto"`
}

func (*AmiTimSpmi) TableName() string {
	return "ami_tim_spmi"
}
