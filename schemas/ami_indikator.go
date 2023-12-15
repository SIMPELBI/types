package model

import "github.com/golang-module/carbon/v2"

type AmiIndikator struct {
	IDIndikator int           `gorm:"primaryKey;column:id_indikator" json:"-"`
	IDAmi       int           `gorm:"column:id_ami" json:"idAmi"`
	Question1   string        `gorm:"column:question1" json:"question1"`
	Question2   string        `gorm:"column:question2" json:"question2"`
	Question3   string        `gorm:"column:question3" json:"question3"`
	Question4   string        `gorm:"column:question4" json:"question4"`
	Question5   string        `gorm:"column:question5" json:"question5"`
	Question6   string        `gorm:"column:question6" json:"question6"`
	Tgl         carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor   int           `gorm:"column:id_auditor" json:"idAuditor"`
}

func (m *AmiIndikator) TableName() string {
	return "ami_indikator"
}
