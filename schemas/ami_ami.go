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

type LaporanAMI struct {
	IDAmi                   int         `gorm:"primaryKey;column:id_ami" json:"-"`
	IDFakultas              int         `gorm:"column:id_fakultas" json:"id_fakultas"`
	Fakultas                string      `gorm:"column:fakultas" json:"fakultas,omitempty"`
	IDProdiUnit             int         `gorm:"column:id_prodi_unit" json:"id_prodi_unit"`
	ProdiUnit               string      `gorm:"column:prodi_unit" json:"prodi_unit,omitempty"`
	NamaKaprodiDir          string      `gorm:"column:nama" json:"nama_kaprodi_dir"`
	TelponKaprodiDir        string      `gorm:"column:telp" json:"telpon_kaprodi_dir"`
	IDJenjang               int         `gorm:"column:id_jenjang" json:"id_jenjang"`
	Jenjang                 string      `gorm:"column:jenjang" json:"jenjang"`
	IDAuditorKetua          int         `gorm:"column:id_auditor_ketua" json:"id_auditor_ketua"`
	NmAuditorKetua          string      `gorm:"column:nm_auditor" json:"nm_auditor_ketua,omitempty"`
	FakultasKetuaAuditor    string      `gorm:"column:fakultas" json:"fakultas_ketua_auditor"`
	TelponAuditorKetua      string      `gorm:"column:telp" json:"telpon_auditor_ketua"`
	IDAnggota1              int         `gorm:"column:id_anggota1" json:"id_anggota1"`
	NmAuditor1              string      `gorm:"column:nm_auditor" json:"nm_auditor_1,omitempty"`
	FakultasAnggotaAuditor1 string      `gorm:"column:fakultas" json:"fakultas_anggota_auditor_1"`
	TelponAnggotaAuditor1   string      `gorm:"column:telp" json:"telpon_anggota_auditor_1"`
	IDAnggota2              int         `gorm:"column:id_anggota2" json:"id_anggota2"`
	NmAuditor2              string      `gorm:"column:nm_auditor" json:"nm_auditor_2,omitempty"`
	FakultasAnggotaAuditor2 string      `gorm:"column:fakultas" json:"fakultas_anggota_auditor_2"`
	TelponAnggotaAuditor2   string      `gorm:"column:telp" json:"telpon_anggota_auditor_2"`
	IDSiklus                int         `gorm:"column:id_siklus" json:"id_siklus"`
	Tahun                   int         `gorm:"column:tahun" json:"tahun"`
	Status                  string      `gorm:"column:status" json:"status"`
	TglRtm                  carbon.Date `gorm:"column:tgl_rtm" json:"tgl_rtm,omitempty"`
	TglSelesai              carbon.Date `gorm:"column:tgl_selesai" json:"tgl_selesai,omitempty"`
	IDStandar               int         `gorm:"column:id_standar" json:"id_standar"`
	Standar                 string      `gorm:"column:standar" json:"standar,omitempty"`
	IsiStandar              string      `gorm:"column:isi" json:"isi_standar"`
	IDIndikator             int         `gorm:"column:id_indikator" json:"id_indikator"`
	NamaIndikator           string      `gorm:"column:nama_indikator" json:"nama_indikator"`
	IsiIndikator            string      `gorm:"column:isi" json:"isi_indikator"`
	Jawaban                 string      `gorm:"column:jawaban" json:"jawaban"`
	IDKts                   int         `gorm:"column:id_kts" json:"id_kts"`
	IDAudit                 int         `gorm:"column:id_audit" json:"id_audit"`
	Kts                     string      `gorm:"column:kts" json:"kts"`
	Uraian                  string      `gorm:"column:uraian" json:"uraian"`
	Tindakan                string      `gorm:"column:tindakan" json:"tindakan"`
	Target                  string      `gorm:"column:target" json:"target"`
	Email                   string      `gorm:"column:email" json:"email"`
}

func (m *AmiAmi) TableName() string {
	return "ami_ami"
}
