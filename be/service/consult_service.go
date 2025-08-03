package service

import (
	"consult-scheduler/model"
	"consult-scheduler/repository"
	"context"
	"time"
)

type ConsultService struct {
	r *repository.ConsultRepository
}

func NewConsultService(r *repository.ConsultRepository) *ConsultService {
	return &ConsultService{r: r}
}

func (s ConsultService) GetDepartmentConsultsForDay(
	ctx context.Context,
	departmentID uint,
	day time.Time,
) ([]*model.ConsultDto, error) {
	consults, err := s.r.FindAllForDepartmentAndDay(ctx, departmentID, day)
	if err != nil {
		return nil, err
	}

	result := make([]*model.ConsultDto, 0, len(consults))
	for _, consult := range consults {
		result = append(result, consult.ToApiModel())
	}

	return result, nil
}

func (s ConsultService) DeleteDepartmentConsult(
	ctx context.Context,
	departmentId uint,
	consultId uint,
) error {
	return s.r.Delete(ctx, departmentId, consultId)
}

func (s ConsultService) CreateDepartmentConsult(
	ctx context.Context,
	departmentID uint,
	consult *model.CreateConsultDto,
) (*model.ConsultDto, error) {
	if consult.DepartmentID != departmentID {
		return nil, model.CustomValidationError{
			FieldName: "departmentId",
			Reason:    "departmentId must match the department ID in path",
		}
	}
	saved, err := s.r.Save(ctx, consult.ToDbModel())
	if err != nil {
		return nil, err
	}
	return saved.ToApiModel(), nil
}

func (s ConsultService) UpdateDepartmentConsult(
	ctx context.Context,
	departmentID uint,
	consultId uint,
	consult *model.CreateConsultDto,
) (*model.ConsultDto, error) {
	if consult.DepartmentID != departmentID {
		return nil, model.CustomValidationError{
			FieldName: "departmentId",
			Reason:    "departmentId must match the department ID in path",
		}
	}
	saved, err := s.r.Update(ctx, consultId, consult.ToDbModel())
	if err != nil {
		return nil, err
	}
	return saved.ToApiModel(), nil
}
