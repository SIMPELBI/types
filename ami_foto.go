package types

import "github.com/golang-module/carbon/v2"

type AmiFoto struct {
	IDFoto    int           `gorm:"primaryKey;column:id_foto" json:"-"`
	IDAmi     int           `gorm:"column:id_ami" json:"idAmi"`
	Foto      string        `gorm:"column:foto" json:"foto"`
	Tgl       carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor int           `gorm:"column:id_auditor" json:"idAuditor"`
}

func (m *AmiFoto) TableName() string {
	return "ami_foto"
}
