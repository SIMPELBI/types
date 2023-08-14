package model

import (
	"database/sql"
	"gorm.io/datatypes"
)

type AmiAmi struct {
	IDAmi          int            `gorm:"primaryKey;column:id_ami" json:"-"`
	IDFakultas     int            `gorm:"column:id_fakultas" json:"idFakultas"`
	IDProdi        int            `gorm:"column:id_prodi" json:"idProdi"`
	IDAuditorKetua int            `gorm:"column:id_auditor_ketua" json:"idAuditorKetua"`
	IDAnggota1     sql.NullInt32  `gorm:"column:id_anggota1" json:"idAnggota1"`
	IDAnggota2     sql.NullInt32  `gorm:"column:id_anggota2" json:"idAnggota2"`
	IDSiklus       int            `gorm:"column:id_siklus" json:"idSiklus"`
	Status         sql.NullString `gorm:"column:status" json:"status"`
	TglRtm         datatypes.Date `gorm:"column:tgl_rtm" json:"tglRtm"`
	TglSelesai     datatypes.Date `gorm:"column:tgl_selesai" json:"tglSelesai"`
}

func (m *AmiAmi) TableName() string {
	return "ami_ami"
}
