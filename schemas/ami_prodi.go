package model

import "database/sql"

type AmiProdi struct {
	IDProdi    int            `gorm:"primaryKey;column:id_prodi" json:"-"`
	IDFakultas int            `gorm:"column:id_fakultas" json:"id_fakultas"`
	UID        string         `gorm:"primaryKey;column:uid" json:"-"`
	Prodi      string         `gorm:"column:prodi" json:"prodi"`
	Jenjang    string         `gorm:"column:jenjang" json:"jenjang"`
	Kaprodi    string         `gorm:"column:kaprodi" json:"kaprodi"`
	Nidn       string         `gorm:"column:nidn" json:"nidn"`
	Password   string         `gorm:"column:password" json:"password"`
	Niknip     sql.NullString `gorm:"column:niknip" json:"niknip"`
	Telp       sql.NullString `gorm:"column:telp" json:"telp"`
	Email      sql.NullString `gorm:"column:email" json:"email"`
	Foto       sql.NullString `gorm:"column:foto" json:"foto"`
	IDSiklus   sql.NullInt32  `gorm:"column:id_siklus" json:"id_siklus"`
}

// TableName get sql table name.获取数据库表名
func (m *AmiProdi) TableName() string {
	return "ami_prodi"
}
