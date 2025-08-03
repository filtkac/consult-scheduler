package repository

import (
	"consult-scheduler/model"
	"context"
	"gorm.io/gorm"
	"time"
)

type ConsultRepository struct {
	db *gorm.DB
}

func NewConsultRepository(db *gorm.DB) *ConsultRepository {
	return &ConsultRepository{db: db}
}

func (r ConsultRepository) FindAllForDepartmentAndDay(
	ctx context.Context,
	departmentID uint,
	day time.Time,
) ([]*model.Consult, error) {
	return gorm.G[*model.Consult](r.db).
		Where("department_id = ?", departmentID).
		Where("time between ? and ?", day, day.AddDate(0, 0, 1)).
		Find(ctx)
}

func (r ConsultRepository) Delete(ctx context.Context, departmentId uint, consultId uint) error {
	_, err := gorm.G[model.Consult](r.db).
		Where("department_id = ?", departmentId).
		Where("id = ?", consultId).
		Delete(ctx)
	return err
}

func (r ConsultRepository) Save(ctx context.Context, consult *model.Consult) (*model.Consult, error) {
	err := gorm.G[model.Consult](r.db).Create(ctx, consult)
	if err != nil {
		return nil, err
	}
	return consult, nil
}

func (r ConsultRepository) Update(ctx context.Context, id uint, consult *model.Consult) (*model.Consult, error) {
	_, err := gorm.G[*model.Consult](r.db).
		Where("id = ?", id).
		Updates(ctx, consult)
	if err != nil {
		return nil, err
	}
	consult.ID = id
	return consult, nil
}
