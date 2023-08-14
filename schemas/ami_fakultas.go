package model

import "database/sql"

type AmiFakultas struct {
	IDFakultas int            `gorm:"primaryKey;column:id_fakultas" json:"-"`
	Fakultas   string         `gorm:"column:fakultas" json:"fakultas"`
	Dekan      string         `gorm:"column:dekan" json:"dekan"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	Password   string         `gorm:"column:password" json:"password"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"idSiklus"`
}

func (m *AmiFakultas) TableName() string {
	return "ami_fakultas"
}
