package service

import (
	"context"
	"github.com/filtkac/consult-scheduler/model"
	"github.com/filtkac/consult-scheduler/repository"
)

type ConsultTypeService struct {
	r *repository.ConsultTypeRepository
}

func NewConsultTypeService(r *repository.ConsultTypeRepository) *ConsultTypeService {
	return &ConsultTypeService{r: r}
}

func (s ConsultTypeService) GetAllConsultTypes(ctx context.Context) ([]*model.ConsultTypeDto, error) {
	consultTypes, err := s.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.ConsultTypeDto, 0, len(consultTypes))
	for _, consultType := range consultTypes {
		result = append(result, consultType.ToApiModel())
	}

	return result, nil
}

func (s ConsultTypeService) DeleteConsultType(ctx context.Context, id uint) error {
	return s.r.Delete(ctx, id)
}

func (s ConsultTypeService) CreateConsultType(
	ctx context.Context,
	consultType *model.ConsultTypeDto,
) (*model.ConsultTypeDto, error) {
	saved, err := s.r.Save(ctx, consultType.ToDbModel())
	if err != nil {
		return nil, err
	}
	return saved.ToApiModel(), nil
}
