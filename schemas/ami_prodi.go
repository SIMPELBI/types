package schemas

type AmiProdi struct {
	IDProdi    int    `gorm:"primaryKey;column:id_prodi" json:"-"`
	IDFakultas int    `gorm:"column:id_fakultas" json:"idFakultas"`
	Prodi      string `gorm:"column:prodi" json:"prodi"`
	Jenjang    string `gorm:"column:jenjang" json:"jenjang"`
	Kaprodi    string `gorm:"column:kaprodi" json:"kaprodi"`
	Nidn       string `gorm:"column:nidn" json:"nidn"`
	Niknip     string `gorm:"column:niknip" json:"niknip"`
	Telp       string `gorm:"column:telp" json:"telp"`
	Email      string `gorm:"column:email" json:"email"`
	Foto       string `gorm:"column:foto" json:"foto"`
	Password   string `gorm:"column:password" json:"password"`
	IDSiklus   int    `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiProdi) TableName() string {
	return "ami_prodi"
}
