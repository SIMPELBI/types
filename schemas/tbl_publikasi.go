package model

type TblPublikasi struct {
	Id              int    `gorm:"primaryKey;column:id" json:"-"`
	NamaDosen       string `gorm:"column:nama_dosen" json:"nama_dosen"`
	KodeProdi       int    `gorm:"column:kode_prodi" json:"kode_prodi"`
	JudulPaper      string `gorm:"column:judul_paper" json:"judul_paper"`
	TahunTerbit     string `gorm:"column:tahun_terbit" json:"tahun_terbit"`
	NamaPublikasi   string `gorm:"column:nama_publikasi" json:"nama_publikasi"`
	JenisPublikasi  string `gorm:"column:jenis_publikasi" json:"jenis_publikasi"`
	Ranking         string `gorm:"column:ranking" json:"ranking"`
	JenisPenelitian string `gorm:"jenis_penelitian" json:"jenis_penelitian"`
	Autor           string `gorm:"autor" json:"autor"`
	JumlahKutipan   int    `gorm:"jumlah_kutipan" json:"jumlah_kutipan"`
}

func (m *TblPublikasi) TableName() string {
	return "publikasi"
}
