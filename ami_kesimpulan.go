package types

import "github.com/golang-module/carbon/v2"

type AmiKesimpulan struct {
	IDKesimpulan int           `gorm:"primaryKey;column:id_kesimpulan" json:"-"`
	IDAmi        int           `gorm:"column:id_ami" json:"idAmi"`
	CkpLengkap   string        `gorm:"column:ckp_lengkap" json:"ckpLengkap"`
	Sebutkan     string        `gorm:"column:sebutkan" json:"sebutkan"`
	Tgl          carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int           `gorm:"column:id_auditor" json:"idAuditor"`
}

func (m *AmiKesimpulan) TableName() string {
	return "ami_kesimpulan"
}
