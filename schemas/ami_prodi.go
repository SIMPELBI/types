package model

type AmiProdi struct {
	IDProdi    int        `gorm:"primaryKey;column:id_prodi" json:"-"`
	UserName   string     `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	IDFakultas int        `gorm:"column:id_fakultas" json:"id_fakultas"`
	Prodi      string     `gorm:"column:prodi" json:"prodi"`
	Jenjang    string     `gorm:"column:jenjang" json:"jenjang"`
	Kaprodi    string     `gorm:"column:kaprodi" json:"kaprodi"`
	Nidn       string     `gorm:"column:nidn" json:"nidn"`
	Niknip     string     `gorm:"column:niknip" json:"niknip"`
	Telp       string     `gorm:"column:telp" json:"telp"`
	Email      string     `gorm:"column:email" json:"email"`
	Foto       string     `gorm:"column:foto" json:"foto"`
	IDSiklus   int        `gorm:"column:id_siklus" json:"id_siklus"`
}

type AmiProdiJoin struct {
	IDProdi    int        `gorm:"primaryKey;column:id_prodi" json:"-"`
	UserName   string     `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	IDFakultas int        `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas   string     `gorm:"column:fakultas" json:"fakultas"`
	Prodi      string     `gorm:"column:prodi" json:"prodi"`
	Jenjang    string     `gorm:"column:jenjang" json:"jenjang"`
	Kaprodi    string     `gorm:"column:kaprodi" json:"kaprodi"`
	Nidn       string     `gorm:"column:nidn" json:"nidn"`
	Niknip     string     `gorm:"column:niknip" json:"niknip"`
	Telp       string     `gorm:"column:telp" json:"telp"`
	Email      string     `gorm:"column:email" json:"email"`
	Foto       string     `gorm:"column:foto" json:"foto"`
	IDSiklus   int        `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiProdi) TableName() string {
	return "ami_prodi"
}
