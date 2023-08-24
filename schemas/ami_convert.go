package model

type AmiConvert struct {
	IDRtm     string `gorm:"primaryKey;column:id_rtm" json:"-"`
	UserLevel int    `gorm:"column:user_level" json:"user_level"`
}

func (m *AmiConvert) TableName() string {
	return "ami_convert"
}
