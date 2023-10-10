package model

type AmiConvert struct {
	IDRtm     string `gorm:"primaryKey;column:id_rtm" json:"-"`
	UserLevel int    `gorm:"column:user_level" json:"user_level"`
}

type AmiConvertProfile struct {
	IDRtm     string `gorm:"primaryKey;column:id_rtm" json:"-"`
	UserLevel int    `gorm:"column:user_level" json:"user_level"`
	NamaLevel string `gorm:"column:nama_level" json:"nama_level"`
	NamaUser  string `gorm:"column:nama_user" json:"nama_user"`
	Email     string `gorm:"column:email" json:"email"`
	Nidn      string `gorm:"column:nidn" json:"nidn"`
	Foto      string `gorm:"column:foto" json:"foto"`
}

func (m *AmiConvert) TableName() string {
	return "ami_convert"
}
