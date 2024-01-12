package model

type TblUserLevel struct {
	IDUserLevel int    `gorm:"primaryKey;column:id_user_level" json:"-"`
	NamaLevel   string `gorm:"column:nama_level" json:"nama_level"`
}

func (m *TblUserLevel) TableName() string {
	return "ami_user_level"
}
