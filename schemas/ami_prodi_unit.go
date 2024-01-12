package model

type AmiProdiUnit struct {
	IDProdiUnit int        `gorm:"primaryKey;column:id_prodi_unit" json:"-"`
	UserName    string     `gorm:"column:user_name" json:"user_name"`
	AmiConvert  AmiConvert `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	IDFakultas  int        `gorm:"column:id_fakultas" json:"id_fakultas"`
	ProdiUnit   string     `gorm:"column:prodi_unit" json:"prodi_unit"`
	IDJenjang   int        `gorm:"column:id_jenjang" json:"id_jenjang"`
	Nama        string     `gorm:"column:nama" json:"nama"`
	Nidn        string     `gorm:"column:nidn" json:"nidn"`
	Niknip      string     `gorm:"column:niknip" json:"niknip"`
	Telp        string     `gorm:"column:telp" json:"telp"`
	Email       string     `gorm:"column:email" json:"email"`
	Foto        string     `gorm:"column:foto" json:"foto"`
	IDSiklus    int        `gorm:"column:id_siklus" json:"id_siklus"`
}

type AmiProdiUnitJoin struct {
	IDProdiUnit int        `gorm:"primaryKey;column:id_prodi_unit" json:"-"`
	UserName    string     `gorm:"column:user_name" json:"user_name"`
	AmiConvert  AmiConvert `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	IDFakultas  int        `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas    string     `gorm:"column:fakultas" json:"fakultas"`
	ProdiUnit   string     `gorm:"column:prodi_unit" json:"prodi_unit"`
	IDJenjang   int        `gorm:"column:id_jenjang" json:"id_jenjang"`
	Jenjang     string     `gorm:"column:jenjang" json:"jenjang"`
	Nama        string     `gorm:"column:nama" json:"nama"`
	Nidn        string     `gorm:"column:nidn" json:"nidn"`
	Niknip      string     `gorm:"column:niknip" json:"niknip"`
	Telp        string     `gorm:"column:telp" json:"telp"`
	Email       string     `gorm:"column:email" json:"email"`
	Foto        string     `gorm:"column:foto" json:"foto"`
	IDSiklus    int        `gorm:"column:id_siklus" json:"id_siklus"`
	Tahun       string     `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiProdiUnit) TableName() string {
	return "ami_prodi_unit"
}
