package model

type Student struct {
	ID          uint64  `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	StudentName string  `json:"name"`
	Classes     []Class `gorm:"many2many:student_classes"`
	Marks       []Mark  `gorm:"foreignkey:StudentID"`
}
