package model

import "database/sql"

type AmiMenu struct {
	IDMenu     int           `gorm:"primaryKey;column:id_menu" json:"-"`
	Title      string        `gorm:"column:title" json:"title"`
	Url        string        `gorm:"column:url" json:"url"`
	Icon       string        `gorm:"column:icon" json:"icon"`
	IsMainMenu int           `gorm:"column:is_main_menu" json:"is_main_menu"`
	IsAktif    string        `gorm:"column:is_aktif" json:"is_aktif"`
	UrutanMenu sql.NullInt32 `gorm:"column:urutan_menu" json:"urutan_menu"`
}

func (m *AmiMenu) TableName() string {
	return "ami_menu"
}
