package model

import (
	"database/sql"
	"github.com/golang-module/carbon/v2"
)

type AmiAmi struct {
	IDAmi          int            `gorm:"primaryKey;column:id_ami" json:"-"`
	IDFakultas     int            `gorm:"column:id_fakultas" json:"id_fakultas"`
	IDProdi        int            `gorm:"column:id_prodi" json:"id_prodi"`
	IDAuditorKetua int            `gorm:"column:id_auditor_ketua" json:"id_auditor_ketua"`
	IDAnggota1     sql.NullInt32  `gorm:"column:id_anggota1" json:"id_anggota1"`
	IDAnggota2     sql.NullInt32  `gorm:"column:id_anggota2" json:"id_anggota2"`
	IDSiklus       int            `gorm:"column:id_siklus" json:"id_siklus"`
	Status         sql.NullString `gorm:"column:status" json:"status"`
	TglRtm         carbon.Date    `gorm:"column:tgl_rtm" json:"tgl_rtm"`
	TglSelesai     carbon.Date    `gorm:"column:tgl_selesai" json:"tgl_selesai"`
}

func (m *AmiAmi) TableName() string {
	return "ami_ami"
}
