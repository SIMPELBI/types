package schemas

type AmiFakultas struct {
	IDFakultas int    `gorm:"primaryKey;column:id_fakultas" json:"-"`
	Fakultas   string `gorm:"column:fakultas" json:"fakultas"`
	Dekan      string `gorm:"column:dekan" json:"dekan"`
	Nidn       string `gorm:"column:nidn" json:"nidn"`
	Niknip     string `gorm:"column:niknip" json:"niknip"`
	Telp       string `gorm:"column:telp" json:"telp"`
	Email      string `gorm:"column:email" json:"email"`
	Foto       string `gorm:"column:foto" json:"foto"`
	Password   string `gorm:"column:password" json:"password"`
	IDSiklus   int    `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiFakultas) TableName() string {
	return "ami_fakultas"
}
