package model

type AmiJenjang struct {
	IDJenjang int    `gorm:"primaryKey;column:id_jenjang" json:"-"`
	Jenjang   string `gorm:"column:jenjang" json:"jenjang"`
}

func (m *AmiJenjang) TableName() string {
	return "ami_jenjang"
}
