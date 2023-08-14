package model

import (
	"database/sql"
	"github.com/golang-module/carbon/v2"
	"gorm.io/datatypes"
)

type AmiAudit struct {
	IDAudit      int            `gorm:"primaryKey;column:id_audit" json:"-"`
	IDAmi        int            `gorm:"column:id_ami" json:"idAmi"`
	IDStandar    int            `gorm:"column:id_standar" json:"idStandar"`
	IDKts        int            `gorm:"column:id_kts" json:"idKts"`
	Uraian       string         `gorm:"column:uraian" json:"uraian"`
	Tindakan     sql.NullString `gorm:"column:tindakan" json:"tindakan"`
	Target       sql.NullString `gorm:"column:target" json:"target"`
	Status       string         `gorm:"column:status" json:"status"`
	Tgl          carbon.Date    `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int            `gorm:"column:id_auditor" json:"idAuditor"`
	TglPerbaikan datatypes.Date `gorm:"column:tgl_perbaikan" json:"tglPerbaikan"`
}

func (m *AmiAudit) TableName() string {
	return "ami_audit"
}
