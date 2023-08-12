package types

type AmiSetting struct {
	IDSetting int    `gorm:"primaryKey;column:id_setting" json:"-"`
	NmPt      string `gorm:"column:nm_pt" json:"nmPt"`
	Footer    string `gorm:"column:footer" json:"footer"`
	Favicon   string `gorm:"column:favicon" json:"favicon"`
	Logo      string `gorm:"column:logo" json:"logo"`
	Kampus    string `gorm:"column:kampus" json:"kampus"`
}

func (m *AmiSetting) TableName() string {
	return "ami_setting"
}
