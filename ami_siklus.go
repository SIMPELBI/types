package types

import "github.com/golang-module/carbon/v2"

type AmiSiklus struct {
	IDSiklus int           `gorm:"primaryKey;column:id_siklus" json:"-"`
	Tahun    carbon.Carbon `gorm:"column:tahun" json:"tahun"`
}

func (m *AmiSiklus) TableName() string {
	return "ami_siklus"
}
