package types

type AmiKts struct {
	IDKts int    `gorm:"primaryKey;column:id_kts" json:"-"`
	Kts   string `gorm:"column:kts" json:"kts"`
}

func (m *AmiKts) TableName() string {
	return "ami_kts"
}
