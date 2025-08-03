package model

type ConsultTypeDto struct {
	ID              uint   `json:"id"`
	ConsultType     string `json:"consultType" binding:"required"`
	DurationMinutes uint   `json:"durationMinutes" binding:"required"`
}

type DepartmentDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

type DayTemplateDto struct {
	ID           uint   `json:"id"`
	Name         string `json:"name" binding:"required"`
	DepartmentID uint   `json:"departmentId" binding:"required"`
}

type DayTemplateConsultDto struct {
	ID            uint     `json:"id"`
	ConsultTypeID uint     `json:"consultTypeId" binding:"required"`
	DayTemplateID uint     `json:"dayTemplateId" binding:"required"`
	Time          TimeOnly `json:"time" binding:"required"`
	Note          *string  `json:"note,omitempty"`
}

type PatientDto struct {
	ID       uint    `json:"id"`
	UniqueID string  `json:"uniqueId" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
}

type ConsultDto struct {
	ID           uint             `json:"id"`
	DepartmentID uint             `json:"departmentId" binding:"required"`
	ConsultType  *ConsultTypeDto  `json:"consultType,omitempty"`
	Time         TimestampMinutes `json:"time" binding:"required"`
	Patient      *PatientDto      `json:"patient,omitempty"`
	Note         *string          `json:"note,omitempty"`
}

type CreateConsultDto struct {
	DepartmentID  uint             `json:"departmentId" binding:"required"`
	ConsultTypeId *uint            `json:"consultTypeId"`
	Time          TimestampMinutes `json:"time" binding:"required"`
	PatientID     *uint            `json:"patient"`
	Note          *string          `json:"note"`
}
