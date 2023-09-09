package model

import "database/sql"

type AmiAuditor struct {
	IDAuditor  int            `gorm:"primaryKey;column:id_auditor" json:"-"`
	UserName   string         `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert     `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	IDFakultas int            `gorm:"column:id_fakultas" json:"id_fakultas"`
	IDProdi    int            `gorm:"column:id_prodi" json:"id_prodi"`
	Auditor    string         `gorm:"column:auditor" json:"auditor"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiAuditor) TableName() string {
	return "ami_auditor"
}
