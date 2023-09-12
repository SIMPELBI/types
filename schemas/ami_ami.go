package model

import (
	"database/sql"
	"time"
)

type AmiAmi struct {
	IDAmi          int            `gorm:"primaryKey;column:id_ami" json:"-"`
	IDFakultas     int            `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas       string         `gorm:"column:fakultas" json:"fakultas"`
	IDProdi        int            `gorm:"column:id_prodi" json:"id_prodi"`
	Prodi          string         `gorm:"column:prodi" json:"prodi"`
	IDAuditorKetua int            `gorm:"column:id_auditor_ketua" json:"id_auditor_ketua"`
	NmAuditorKetua string         `gorm:"column:nm_auditor_ketua" json:"nm_auditor_ketua"`
	IDAnggota1     sql.NullInt32  `gorm:"column:id_anggota1" json:"id_anggota1"`
	NmAuditor1     string         `gorm:"column:nm_auditor1" json:"nm_auditor1"`
	IDAnggota2     sql.NullInt32  `gorm:"column:id_anggota2" json:"id_anggota2"`
	NmAuditor2     string         `gorm:"column:nm_auditor2" json:"nm_auditor2"`
	IDSiklus       int            `gorm:"column:id_siklus" json:"id_siklus"`
	Tahun          int            `gorm:"column:tahun" json:"tahun"`
	Status         sql.NullString `gorm:"column:status" json:"status"`
	TglRtm         time.Time      `gorm:"column:tgl_rtm" json:"tgl_rtm"`
	TglSelesai     sql.NullTime   `gorm:"column:tgl_selesai" json:"tgl_selesai"`
}

func (m *AmiAmi) TableName() string {
	return "ami_ami"
}
