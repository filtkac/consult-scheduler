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
