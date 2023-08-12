package types

type AmiAuditor struct {
	IDAuditor  int    `gorm:"primaryKey;column:id_auditor" json:"-"`
	IDFakultas int    `gorm:"column:id_fakultas" json:"idFakultas"`
	IDProdi    int    `gorm:"column:id_prodi" json:"idProdi"`
	Auditor    string `gorm:"column:auditor" json:"auditor"`
	Nidn       string `gorm:"column:nidn" json:"nidn"`
	Niknip     string `gorm:"column:niknip" json:"niknip"`
	Telp       string `gorm:"column:telp" json:"telp"`
	Email      string `gorm:"column:email" json:"email"`
	Foto       string `gorm:"column:foto" json:"foto"`
	Password   string `gorm:"column:password" json:"password"`
	IDSiklus   int    `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiAuditor) TableName() string {
	return "ami_auditor"
}
