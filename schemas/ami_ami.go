package model

import (
	"github.com/golang-module/carbon/v2"
)

type AmiAmi struct {
	IDAmi          int           `gorm:"primaryKey;column:id_ami" json:"-"`
	IDFakultas     int           `gorm:"column:id_fakultas" json:"id_fakultas"`
	IDProdiUnit    int           `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
	IDAuditorKetua int           `gorm:"column:id_auditor_ketua" json:"id_auditor_ketua"`
	IDStandar      int           `gorm:"column:id_standar" json:"id_standar"`
	IDAnggota1     int           `gorm:"column:id_anggota1" json:"id_anggota1"`
	IDAnggota2     int           `gorm:"column:id_anggota2" json:"id_anggota2"`
	IDSiklus       int           `gorm:"column:id_siklus" json:"id_siklus"`
	Status         string        `gorm:"column:status" json:"status"`
	TglRtm         carbon.Carbon `gorm:"column:tgl_rtm" json:"tgl_rtm"`
	TglSelesai     carbon.Carbon `gorm:"column:tgl_selesai" json:"tgl_selesai"`
}

type AmiAmiJoin struct {
	IDAmi          int         `gorm:"primaryKey;column:id_ami" json:"-"`
	IDFakultas     int         `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas       string      `gorm:"column:fakultas" json:"fakultas"`
	IDProdiUnit    int         `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
	ProdiUnit      string      `gorm:"column:prodi_unit" json:"prodi_unit"`
	IDAuditorKetua int         `gorm:"column:id_auditor_ketua" json:"id_auditor_ketua"`
	NmAuditorKetua string      `gorm:"column:nm_auditor_ketua" json:"nm_auditor_ketua"`
	IDAnggota1     int         `gorm:"column:id_anggota1" json:"id_anggota1"`
	NmAuditor1     string      `gorm:"column:anggota1" json:"anggota1"`
	IDAnggota2     int         `gorm:"column:id_anggota2" json:"id_anggota2"`
	NmAuditor2     string      `gorm:"column:anggota2" json:"anggota2"`
	IDSiklus       int         `gorm:"column:id_siklus" json:"id_siklus"`
	Tahun          int         `gorm:"column:tahun" json:"tahun"`
	IDStandar      int         `gorm:"column:id_standar" json:"id_standar"`
	Standar        string      `gorm:"column:standar" json:"standar"`
	Status         string      `gorm:"column:status" json:"status"`
	TglRtm         carbon.Date `gorm:"column:tgl_rtm" json:"tgl_rtm"`
	TglSelesai     carbon.Date `gorm:"column:tgl_selesai" json:"tgl_selesai"`
}

type RekapTemuan struct {
	IDAmi      int    `gorm:"column:id_ami" json:"id_ami"`
	IDFakultas int    `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas   string `gorm:"column:fakultas" json:"fakultas"`
	Prodi      string `gorm:"column:prodi" json:"prodi"`
	Observasi  string `gorm:"column:observasi" json:"observasi"`
	Minor      string `gorm:"column:minor" json:"minor"`
	Mayor      string `gorm:"column:mayor" json:"mayor"`
	IDSiklus   int    `gorm:"column:id_siklus" json:"id_siklus"`
	Tahun      string `gorm:"column:tahun" json:"tahun"`
}

type FileLaporanAmi struct {
	IDAmi          int    `gorm:"column:id_ami" json:"id_ami"`
	IDFakultas     int    `gorm:"column:id_fakultas" json:"id_fakultas"`
	IDProdi        int    `gorm:"column:id_prodi" json:"id_prodi"`
	IDAuditorKetua int    `gorm:"column:id_auditor_ketua" json:"id_auditor_ketua"`
	IDAnggota1     int    `gorm:"column:id_anggota1" json:"id_anggota1"`
	IDAnggota2     int    `gorm:"column:id_anggota2" json:"id_anggota2"`
	IDSiklus       int    `gorm:"column:id_siklus" json:"id_siklus"`
	Status         string `gorm:"column:status" json:"status"`
	TglRtm         string `gorm:"column:tgl_rtm" json:"tgl_rtm"`
	TglSelesai     string `gorm:"column:tgl_selesai" json:"tgl_selesai"`
	NmAuditorKetua string `gorm:"column:nm_auditor_ketua" json:"nm_auditor_ketua"`
	NmAuditor1     string `gorm:"column:anggota1" json:"anggota1"`
	NmAuditor2     string `gorm:"column:anggota2" json:"anggota2"`
	Fakultas       string `gorm:"column:fakultas" json:"fakultas"`
	Prodi          string `gorm:"column:prodi" json:"prodi"`
	Tahun          string `gorm:"column:tahun" json:"tahun"`
	IDJenjang      int    `gorm:"column:id_jenjang" json:"id_jenjang"`
	Jenjang        string `gorm:"column:jenjang" json:"jenjang"`
	IDMekanisme    int    `gorm:"column:id_mekanisme" json:"id_mekanisme"`
	Question1      string `gorm:"column:question1" json:"question1"`
	Question2      string `gorm:"column:question2" json:"question2"`
	Question3      string `gorm:"column:question3" json:"question3"`
	Question4      string `gorm:"column:question4" json:"question4"`
	Question5      string `gorm:"column:question5" json:"question5"`
	Question6      string `gorm:"column:question6" json:"question6"`
	IDAudit        int    `gorm:"column:id_audit" json:"id_audit"`
	Uraian         string `gorm:"column:uraian" json:"uraian"`
	IDKesimpulan   int    `gorm:"column:id_kesimpulan" json:"id_kesimpulan"`
	CkpLengkap     string `gorm:"column:ckp_lurator" json:"ckp_lurator"`
	IDKts          int    `gorm:"column:id_kts" json:"id_kts"`
	Kts            string `gorm:"column:kts" json:"kts"`
}

func (m *AmiAmi) TableName() string {
	return "ami_ami"
}
