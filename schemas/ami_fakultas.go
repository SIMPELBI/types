package model

type AmiFakultas struct {
	IDFakultas int        `gorm:"primaryKey;column:id_fakultas" json:"-"`
	UserName   string     `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	Fakultas   string     `gorm:"column:fakultas" json:"fakultas"`
	Dekan      string     `gorm:"column:dekan" json:"dekan"`
	Nidn       string     `gorm:"column:nidn" json:"nidn"`
	Niknip     string     `gorm:"column:niknip" json:"niknip"`
	Telp       string     `gorm:"column:telp" json:"telp"`
	Email      string     `gorm:"column:email" json:"email"`
	Foto       string     `gorm:"column:foto" json:"foto"`
	IDSiklus   int        `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiFakultas) TableName() string {
	return "ami_fakultas"
}
