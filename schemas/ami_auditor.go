package model

import "database/sql"

type AmiAuditor struct {
	IDAuditor  int            `gorm:"primaryKey;column:id_auditor" json:"-"`
	IDFakultas int            `gorm:"column:id_fakultas" json:"idFakultas"`
	IDProdi    int            `gorm:"column:id_prodi" json:"idProdi"`
	Auditor    string         `gorm:"column:auditor" json:"auditor"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	Password   string         `gorm:"column:password" json:"password"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiAuditor) TableName() string {
	return "ami_auditor"
}
