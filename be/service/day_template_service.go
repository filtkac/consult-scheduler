package service

import (
	"context"
	"github.com/filtkac/consult-scheduler/model"
	"github.com/filtkac/consult-scheduler/repository"
)

type DayTemplateService struct {
	r *repository.DayTemplateRepository
}

func NewDayTemplateService(r *repository.DayTemplateRepository) *DayTemplateService {
	return &DayTemplateService{r: r}
}

func (s DayTemplateService) GetAllDayTemplates(ctx context.Context) ([]*model.DayTemplateDto, error) {
	dayTemplates, err := s.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapDayTemplates(dayTemplates), nil
}

func (s DayTemplateService) GetDayTemplatesForDepartment(
	ctx context.Context,
	departmentID uint,
) ([]*model.DayTemplateDto, error) {
	dayTemplates, err := s.r.FindAllForDepartment(ctx, departmentID)
	if err != nil {
		return nil, err
	}
	return mapDayTemplates(dayTemplates), nil
}

func (s DayTemplateService) DeleteDayTemplate(ctx context.Context, id uint) error {
	return s.r.Delete(ctx, id)
}

func (s DayTemplateService) CreateDayTemplate(
	ctx context.Context,
	dayTemplate *model.DayTemplateDto,
) (*model.DayTemplateDto, error) {
	saved, err := s.r.Save(ctx, dayTemplate.ToDbModel())
	if err != nil {
		return nil, err
	}
	return saved.ToApiModel(), nil
}

func (s DayTemplateService) GetConsultsForDayTemplate(
	ctx context.Context,
	dayTemplateID uint,
) ([]*model.DayTemplateConsultDto, error) {
	consults, err := s.r.FindAllConsultsForDayTemplate(ctx, dayTemplateID)
	if err != nil {
		return nil, err
	}

	result := make([]*model.DayTemplateConsultDto, 0, len(consults))
	for _, consult := range consults {
		result = append(result, consult.ToApiModel())
	}

	return result, nil
}

func (s DayTemplateService) CreateConsultsForDayTemplate(
	ctx context.Context,
	dayTemplateID uint,
	consults []*model.DayTemplateConsultDto,
) ([]*model.DayTemplateConsultDto, error) {
	for _, consult := range consults {
		if consult.DayTemplateID != dayTemplateID {
			return nil, model.CustomValidationError{
				FieldName: "dayTemplateId",
				Reason:    "dayTemplateId must match the day template ID in path",
			}
		}
	}

	err := s.r.DeleteAllConsultsForDayTemplate(ctx, dayTemplateID)
	if err != nil {
		return nil, err
	}

	toSave := make([]*model.DayTemplateConsult, 0, len(consults))
	for _, consult := range consults {
		toSave = append(toSave, consult.ToDbModel())
	}
	saved, err := s.r.SaveDayTemplateConsults(ctx, toSave)
	if err != nil {
		return nil, err
	}

	result := make([]*model.DayTemplateConsultDto, 0, len(saved))
	for _, consult := range saved {
		result = append(result, consult.ToApiModel())
	}
	return result, nil
}

func mapDayTemplates(dayTemplates []*model.DayTemplate) []*model.DayTemplateDto {
	result := make([]*model.DayTemplateDto, 0, len(dayTemplates))
	for _, department := range dayTemplates {
		result = append(result, department.ToApiModel())
	}
	return result
}
