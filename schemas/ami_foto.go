package model

import "github.com/golang-module/carbon/v2"

type AmiFoto struct {
	IDFoto    int           `gorm:"primaryKey;column:id_foto" json:"-"`
	IDAmi     int           `gorm:"column:id_ami" json:"id_ami"`
	Foto      string        `gorm:"column:foto" json:"foto"`
	Tgl       carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor int           `gorm:"column:id_auditor" json:"id_auditor"`
}

type AmiFotoJoin struct {
	IDFoto     int           `gorm:"primaryKey;column:id_foto" json:"-"`
	IDAmi      int           `gorm:"column:id_ami" json:"id_ami"`
	IDFakultas int           `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas   string        `gorm:"column:fakultas" json:"fakultas"`
	IDProdi    int           `gorm:"column:id_prodi" json:"id_prodi"`
	Prodi      string        `gorm:"column:prodi" json:"prodi"`
	Foto       string        `gorm:"column:foto" json:"foto"`
	Tgl        carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor  int           `gorm:"column:id_auditor" json:"id_auditor"`
	Auditor    string        `gorm:"column:auditor" json:"auditor"`
}

func (m *AmiFoto) TableName() string {
	return "ami_foto"
}
