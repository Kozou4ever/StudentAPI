package model

type Mark struct {
	ID        uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Value     float64
	StudentID uint
	ClassID   uint
}
