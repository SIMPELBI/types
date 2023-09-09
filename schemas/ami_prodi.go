package model

import "database/sql"

type AmiProdi struct {
	IDProdi    int            `gorm:"primaryKey;column:id_prodi" json:"-"`
	UserName   string         `gorm:"column:user_name" json:"user_name"`
	AmiConvert AmiConvert     `gorm:"joinForeignKey:user_name;foreignKey:id_rtm;references:UserName" json:"ami_convert_list"`
	IDFakultas int            `gorm:"column:id_fakultas" json:"id_fakultas"`
	Prodi      string         `gorm:"column:prodi" json:"prodi"`
	Jenjang    string         `gorm:"column:jenjang" json:"jenjang"`
	Kaprodi    string         `gorm:"column:kaprodi" json:"kaprodi"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiProdi) TableName() string {
	return "ami_prodi"
}
