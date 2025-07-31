package model

type Model struct {
	ID uint `gorm:"primarykey"`
}

type Department struct {
	Model
	Name string `gorm:"unique;not null"`
}

type ConsultType struct {
	Model
	ConsultType     string `gorm:"unique;not null"`
	DurationMinutes uint
}

type DayTemplate struct {
	Model
	Name         string `gorm:"not null"`
	DepartmentID uint   `gorm:"not null"`
	Department   Department
}

type DayTemplateConsult struct {
	Model
	ConsultTypeID uint
	ConsultType   ConsultType
	DayTemplateID uint        `gorm:"not null"`
	DayTemplate   DayTemplate `gorm:"constraint:OnDelete:CASCADE"`
	Time          TimeOnly    `gorm:"not null"`
	Note          *string
}

type Patient struct {
	Model
	UniqueID string `gorm:"unique;not null"`
	Name     string
	Email    string
	Phone    string
}
