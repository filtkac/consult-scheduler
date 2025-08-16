package model

import "time"

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
	Name         string     `gorm:"not null"`
	DepartmentID uint       `gorm:"not null"`
	Department   Department `gorm:"constraint:OnDelete:CASCADE"`
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
	Name     string `gorm:"not null"`
	Email    *string
	Phone    *string
}

type Consult struct {
	Model
	DepartmentID  uint       `gorm:"not null"`
	Department    Department `gorm:"constraint:OnDelete:CASCADE"`
	ConsultTypeID *uint
	ConsultType   ConsultType
	Time          time.Time `gorm:"type:timestamp;not null"`
	Note          *string
	PatientID     *uint
	Patient       Patient
}
