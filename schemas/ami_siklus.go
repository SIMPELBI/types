package model

type AmiSiklus struct {
	IDSiklus int `gorm:"primaryKey;column:id_siklus" json:"-"`
	Tahun    int `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiSiklus) TableName() string {
	return "ami_siklus"
}
