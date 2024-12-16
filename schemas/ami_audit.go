package model

import (
	"github.com/golang-module/carbon/v2"
)

type AmiAudit struct {
	IDAudit       int           `gorm:"primaryKey;column:id_audit" json:"-"`
	IDAmi         int           `gorm:"column:id_ami" json:"id_ami"`
	IDStandar     int           `gorm:"column:id_standar" json:"id_standar"`
	IDKts         int           `gorm:"column:id_kts" json:"id_kts"`
	Uraian        string        `gorm:"column:uraian" json:"uraian"`
	Tindakan      string        `gorm:"column:tindakan" json:"tindakan"`
	Target        string        `gorm:"column:target" json:"target"`
	Status        string        `gorm:"column:status" json:"status"`
	Tgl           carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor     int           `gorm:"column:id_auditor" json:"id_auditor"`
	TglPerbaikan  carbon.Carbon `gorm:"column:tgl_perbaikan" json:"tgl_perbaikan"`
	Jawaban       string        `gorm:"column:jawaban" json:"jawaban"`
	LinkPerbaikan string        `gorm:"column:link_perbaikan" json:"link_perbaikan"`
}

type AmiAuditJoin struct {
	IDAudit       int           `gorm:"primaryKey;column:id_audit" json:"-"`
	IDAmi         int           `gorm:"column:id_ami" json:"id_ami"`
	IDStandar     int           `gorm:"column:id_standar" json:"id_standar"`
	Standar       string        `gorm:"column:standar" json:"standar"`
	IsiStandar    string        `gorm:"column:isi" json:"isi_standar"`
	IDIndikator   int           `gorm:"column:id_indikator" json:"id_indikator"`
	NamaIndikator string        `gorm:"column:nama_indikator" json:"nama_indikator"`
	IsiIndikator  string        `gorm:"column:isi" json:"isi_indikator"`
	IDKts         int           `gorm:"column:id_kts" json:"id_kts"`
	Kts           string        `gorm:"column:kts" json:"kts"`
	Uraian        string        `gorm:"column:uraian" json:"uraian"`
	Tindakan      string        `gorm:"column:tindakan" json:"tindakan"`
	Target        string        `gorm:"column:target" json:"target"`
	Status        string        `gorm:"column:status" json:"status"`
	Tgl           carbon.Carbon `gorm:"column:tgl" json:"tgl"`
	IDAuditor     int           `gorm:"column:id_auditor" json:"id_auditor"`
	Auditor       string        `gorm:"column:auditor" json:"auditor"`
	TglPerbaikan  carbon.Carbon `gorm:"column:tgl_perbaikan" json:"tgl_perbaikan"`
	Jawaban       string        `gorm:"column:jawaban" json:"jawaban"`
	LinkPerbaikan *string        `gorm:"column:link_perbaikan" json:"link_perbaikan"`
}

func (m *AmiAudit) TableName() string {
	return "ami_audit"
}
