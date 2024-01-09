package model

type LogsQuestion struct {
	IDQuestion int    `gorm:"primaryKey;column:id_question" json:"-"`
	IDAmi      int    `gorm:"column:id_ami" json:"id_ami"`
	IDStandar  int    `gorm:"column:id_standar" json:"id_standar"`
	Jawaban    string `gorm:"column:jawaban" json:"jawaban"`
}

func (m *LogsQuestion) TableName() string {
	return "logs_question"
}
