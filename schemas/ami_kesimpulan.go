package model

import (
	"database/sql"
	"github.com/golang-module/carbon/v2"
)

type AmiKesimpulan struct {
	IDKesimpulan int            `gorm:"primaryKey;column:id_kesimpulan" json:"-"`
	IDAmi        int            `gorm:"column:id_ami" json:"id_ami"`
	CkpLengkap   string         `gorm:"column:ckp_lengkap" json:"ckp_lengkap"`
	Sebutkan     sql.NullString `gorm:"column:sebutkan" json:"sebutkan"`
	Tgl          carbon.Carbon  `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int            `gorm:"column:id_auditor" json:"id_auditor"`
}

func (m *AmiKesimpulan) TableName() string {
	return "ami_kesimpulan"
}
