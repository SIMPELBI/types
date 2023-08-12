package schemas

type AmiAdmin struct {
	IDAdmin  int    `gorm:"primaryKey;column:id_admin" json:"-"`
	NmAdmin  string `gorm:"column:nm_admin" json:"nmAdmin"`
	Jabatan  string `gorm:"column:jabatan" json:"jabatan"`
	Email    string `gorm:"column:email" json:"email"`
	Nidn     string `gorm:"column:nidn" json:"nidn"`
	Password string `gorm:"column:password" json:"password"`
	Foto     string `gorm:"column:foto" json:"foto"`
	IDSiklus int    `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiAdmin) TableName() string {
	return "ami_admin"
}
