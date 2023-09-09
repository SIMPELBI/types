package model

import "database/sql"

type AmiFakultas struct {
	IDFakultas int            `gorm:"primaryKey;column:id_fakultas" json:"-"`
	UserName   string         `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert     `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	Fakultas   string         `gorm:"column:fakultas" json:"fakultas"`
	Dekan      string         `gorm:"column:dekan" json:"dekan"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiFakultas) TableName() string {
	return "ami_fakultas"
}
