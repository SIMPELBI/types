package model

import "database/sql"

type TblUserLevel struct {
	IDUserLevel int            `gorm:"primaryKey;column:id_user_level" json:"-"`
	NamaLevel   sql.NullString `gorm:"column:nama_level" json:"nama_level"`
}

func (m *TblUserLevel) TableName() string {
	return "tbl_user_level"
}
