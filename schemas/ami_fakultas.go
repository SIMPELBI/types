package model

import "database/sql"

type AmiFakultas struct {
	IDFakultas int            `gorm:"primaryKey;column:id_fakultas" json:"-"`
	UID        string         `gorm:"primaryKey;column:uid" json:"-"`
	Fakultas   string         `gorm:"column:fakultas" json:"fakultas"`
	Dekan      string         `gorm:"column:dekan" json:"dekan"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Password   string         `gorm:"column:password" json:"password"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

func (m *AmiFakultas) TableName() string {
	return "ami_fakultas"
}
