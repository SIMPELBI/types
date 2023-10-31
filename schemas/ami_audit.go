package model

import (
	"github.com/golang-module/carbon/v2"
)

type AmiAudit struct {
	IDAudit      int           `gorm:"primaryKey;column:id_audit" json:"-"`
	IDAmi        int           `gorm:"column:id_ami" json:"id_ami"`
	IDStandar    int           `gorm:"column:id_standar" json:"id_standar"`
	IDKts        int           `gorm:"column:id_kts" json:"id_kts"`
	Uraian       string        `gorm:"column:uraian" json:"uraian"`
	Tindakan     string        `gorm:"column:tindakan" json:"tindakan"`
	Target       string        `gorm:"column:target" json:"target"`
	Status       string        `gorm:"column:status" json:"status"`
	Tgl          carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int           `gorm:"column:id_auditor" json:"id_auditor"`
	TglPerbaikan carbon.Carbon `gorm:"column:tgl_perbaikan" json:"tgl_perbaikan"`
}

type AmiAuditJoin struct {
	IDAudit      int           `gorm:"primaryKey;column:id_audit" json:"-"`
	IDAmi        int           `gorm:"column:id_ami" json:"id_ami"`
	IDStandar    int           `gorm:"column:id_standar" json:"id_standar"`
	Standar      string        `gorm:"column:standar" json:"standar"`
	UntukPilihan string        `gorm:"column:utk_pilihan" json:"utk_pilihan"`
	Isi          string        `gorm:"column:isi" json:"isi"`
	IDKts        int           `gorm:"column:id_kts" json:"id_kts"`
	Kts          string        `gorm:"column:kts" json:"kts"`
	Uraian       string        `gorm:"column:uraian" json:"uraian"`
	Tindakan     string        `gorm:"column:tindakan" json:"tindakan"`
	Target       string        `gorm:"column:target" json:"target"`
	Status       string        `gorm:"column:status" json:"status"`
	Tgl          carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor    int           `gorm:"column:id_auditor" json:"id_auditor"`
	Auditor      string        `gorm:"column:auditor" json:"auditor"`
	TglPerbaikan carbon.Carbon `gorm:"column:tgl_perbaikan" json:"tgl_perbaikan"`
}

func (m *AmiAudit) TableName() string {
	return "ami_audit"
}
