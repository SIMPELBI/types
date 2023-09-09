package model

import (
	"database/sql"
	"github.com/golang-module/carbon/v2"
)

type AmiAudit struct {
	IDAudit      int            `gorm:"primaryKey;column:id_audit" json:"-"`
	IDAmi        int            `gorm:"column:id_ami" json:"id_ami"`
	IDStandar    int            `gorm:"column:id_standar" json:"id_standar"`
	IDKts        int            `gorm:"column:id_kts" json:"id_kts"`
	Uraian       string         `gorm:"column:uraian" json:"uraian"`
	Tindakan     sql.NullString `gorm:"column:tindakan" json:"tindakan"`
	Target       sql.NullString `gorm:"column:target" json:"target"`
	Status       string         `gorm:"column:status" json:"status"`
	Tgl          carbon.Carbon  `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int            `gorm:"column:id_auditor" json:"id_auditor"`
	TglPerbaikan sql.NullTime   `gorm:"column:tgl_perbaikan" json:"tgl_perbaikan"`
}

func (m *AmiAudit) TableName() string {
	return "ami_audit"
}
