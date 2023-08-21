package model

import "database/sql"

type AmiAuditor struct {
	IDAuditor  int            `gorm:"primaryKey;column:id_auditor" json:"-"`
	UID        string         `gorm:"primaryKey;column:uid" json:"-"`
	IDFakultas int            `gorm:"column:id_fakultas" json:"id_fakultas"`
	IDProdi    int            `gorm:"column:id_prodi" json:"id_prodi"`
	Auditor    string         `gorm:"column:auditor" json:"auditor"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Password   string         `gorm:"column:password" json:"password"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiAuditor) TableName() string {
	return "ami_auditor"
}
