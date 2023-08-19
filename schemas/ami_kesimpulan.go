package model

import (
	"database/sql"
)

type AmiKesimpulan struct {
	IDKesimpulan int            `gorm:"primaryKey;column:id_kesimpulan" json:"-"`
	IDAmi        int            `gorm:"column:id_ami" json:"idAmi"`
	CkpLengkap   string         `gorm:"column:ckp_lengkap" json:"ckpLengkap"`
	Sebutkan     sql.NullString `gorm:"column:sebutkan" json:"sebutkan"`
	Tgl          sql.NullTime   `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int            `gorm:"column:id_auditor" json:"idAuditor"`
}

func (m *AmiKesimpulan) TableName() string {
	return "ami_kesimpulan"
}
