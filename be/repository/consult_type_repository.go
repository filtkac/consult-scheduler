package repository

import (
	"context"
	"github.com/filtkac/consult-scheduler/model"
	"gorm.io/gorm"
)

type ConsultTypeRepository struct {
	db *gorm.DB
}

func NewConsultTypeRepository(db *gorm.DB) *ConsultTypeRepository {
	return &ConsultTypeRepository{db: db}
}

func (r ConsultTypeRepository) FindAll(ctx context.Context) ([]*model.ConsultType, error) {
	return gorm.G[*model.ConsultType](r.db).Find(ctx)
}

func (r ConsultTypeRepository) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[*model.ConsultType](r.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (r ConsultTypeRepository) Save(ctx context.Context, consultType *model.ConsultType) (*model.ConsultType, error) {
	err := gorm.G[model.ConsultType](r.db).Create(ctx, consultType)
	if err != nil {
		return nil, err
	}
	return consultType, nil
}
