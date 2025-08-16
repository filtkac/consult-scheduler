package model

import "time"

func (ct *ConsultType) ToApiModel() *ConsultTypeDto {
	return &ConsultTypeDto{
		&ct.ID,
		ct.ConsultType,
		ct.DurationMinutes,
	}
}

func (ct *ConsultTypeDto) ToDbModel() *ConsultType {
	return &ConsultType{
		ConsultType:     ct.ConsultType,
		DurationMinutes: ct.DurationMinutes,
	}
}

func (d *Department) ToApiModel() *DepartmentDto {
	return &DepartmentDto{
		&d.ID,
		d.Name,
	}
}

func (d *DepartmentDto) ToDbModel() *Department {
	return &Department{Name: d.Name}
}

func (dt *DayTemplate) ToApiModel() *DayTemplateDto {
	return &DayTemplateDto{
		&dt.ID,
		dt.Name,
		dt.DepartmentID,
	}
}

func (dtc *DayTemplateConsult) ToApiModel() *DayTemplateConsultDto {
	return &DayTemplateConsultDto{
		&dtc.ID,
		dtc.ConsultTypeID,
		dtc.DayTemplateID,
		dtc.Time,
		dtc.Note,
	}
}

func (dt *DayTemplateDto) ToDbModel() *DayTemplate {
	return &DayTemplate{
		Name:         dt.Name,
		DepartmentID: dt.DepartmentID,
	}
}

func (dtc *DayTemplateConsultDto) ToDbModel() *DayTemplateConsult {
	return &DayTemplateConsult{
		ConsultTypeID: dtc.ConsultTypeID,
		DayTemplateID: dtc.DayTemplateID,
		Time:          dtc.Time,
		Note:          dtc.Note,
	}
}

func (p *Patient) ToApiModel() *PatientDto {
	return &PatientDto{
		&p.ID,
		p.UniqueID,
		p.Name,
		p.Email,
		p.Phone,
	}
}

func (p *PatientDto) ToDbModel() *Patient {
	return &Patient{
		UniqueID: p.UniqueID,
		Name:     p.Name,
		Email:    p.Email,
		Phone:    p.Phone,
	}
}

func (c *Consult) ToApiModel() *ConsultDto {
	dto := &ConsultDto{
		ID:           &c.ID,
		DepartmentID: c.DepartmentID,
		Time:         TimestampMinutes(c.Time),
		Note:         c.Note,
	}

	if c.ConsultTypeID != nil {
		dto.ConsultType = c.ConsultType.ToApiModel()
	}

	if c.PatientID != nil {
		dto.Patient = c.Patient.ToApiModel()
	}

	return dto
}

func (c *CreateConsultDto) ToDbModel() *Consult {
	return &Consult{
		DepartmentID:  c.DepartmentID,
		ConsultTypeID: c.ConsultTypeID,
		Time:          time.Time(c.Time),
		Note:          c.Note,
		PatientID:     c.PatientID,
	}
}
