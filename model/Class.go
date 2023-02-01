package model

type Class struct {
	ID        uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ClassName string    `json:"name"`
	Students  []Student `gorm:"many2many:student_classes;"`
	Marks     []Mark    `gorm:"foreignkey:ClassID"`
}
