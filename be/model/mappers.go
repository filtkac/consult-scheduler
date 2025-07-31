package model

func (ct *ConsultType) ToApiModel() *ConsultTypeDto {
	return &ConsultTypeDto{
		ct.ID,
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
		d.ID,
		d.Name,
	}
}

func (d *DepartmentDto) ToDbModel() *Department {
	return &Department{Name: d.Name}
}

func (dt *DayTemplate) ToApiModel() *DayTemplateDto {
	return &DayTemplateDto{
		dt.ID,
		dt.Name,
		dt.DepartmentID,
	}
}

func (dtc *DayTemplateConsult) ToApiModel() *DayTemplateConsultDto {
	return &DayTemplateConsultDto{
		dtc.ID,
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
