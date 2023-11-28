package model

type AmiStrukturDirektoratPsms struct {
	IDAnggota int    `gorm:"primaryKey;column:id_anggota" json:"-"`
	Nama      string `gorm:"column:nama" json:"nama"`
	Jabatan   string `gorm:"column:jabatan" json:"jabatan"`
	Foto      string `gorm:"column:foto" json:"foto"`
}

func (*AmiStrukturDirektoratPsms) TableName() string {
	return "ami_struktur_direktorat_psms"
}
