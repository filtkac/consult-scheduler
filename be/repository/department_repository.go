package repository

import (
	"consult-scheduler/model"
	"context"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r DepartmentRepository) FindAll(ctx context.Context) ([]*model.Department, error) {
	return gorm.G[*model.Department](r.db).Find(ctx)
}

func (r DepartmentRepository) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[model.Department](r.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (r DepartmentRepository) Save(ctx context.Context, department *model.Department) (*model.Department, error) {
	err := gorm.G[model.Department](r.db).Create(ctx, department)
	if err != nil {
		return nil, err
	}
	return department, nil
}
