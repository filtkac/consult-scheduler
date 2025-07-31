package service

import (
	"consult-scheduler/model"
	"consult-scheduler/repository"
	"context"
)

type DepartmentService struct {
	r *repository.DepartmentRepository
}

func NewDepartmentService(r *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{r: r}
}

func (s DepartmentService) GetAllDepartments(ctx context.Context) ([]*model.DepartmentDto, error) {
	departments, err := s.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.DepartmentDto, 0, len(departments))
	for _, department := range departments {
		result = append(result, department.ToApiModel())
	}

	return result, nil
}

func (s DepartmentService) DeleteDepartment(ctx context.Context, id uint) error {
	return s.r.Delete(ctx, id)
}

func (s DepartmentService) CreateDepartment(
	ctx context.Context,
	department *model.DepartmentDto,
) (*model.DepartmentDto, error) {
	saved, err := s.r.Save(ctx, department.ToDbModel())
	if err != nil {
		return nil, err
	}
	return saved.ToApiModel(), nil
}
