package repository

import (
	"context"
	"github.com/filtkac/consult-scheduler/model"
	"gorm.io/gorm"
)

type DayTemplateRepository struct {
	db *gorm.DB
}

func NewDayTemplateRepository(db *gorm.DB) *DayTemplateRepository {
	return &DayTemplateRepository{db: db}
}

func (r DayTemplateRepository) FindAll(ctx context.Context) ([]*model.DayTemplate, error) {
	return gorm.G[*model.DayTemplate](r.db).Find(ctx)
}

func (r DayTemplateRepository) FindAllForDepartment(
	ctx context.Context,
	departmentID uint,
) ([]*model.DayTemplate, error) {
	return gorm.G[*model.DayTemplate](r.db).
		Where("department_id = ?", departmentID).
		Find(ctx)
}

func (r DayTemplateRepository) Delete(ctx context.Context, id uint) error {
	_, err := gorm.G[model.DayTemplate](r.db).Where("id = ?", id).Delete(ctx)
	return err
}

func (r DayTemplateRepository) Save(ctx context.Context, dayTemplate *model.DayTemplate) (*model.DayTemplate, error) {
	err := gorm.G[model.DayTemplate](r.db).Create(ctx, dayTemplate)
	if err != nil {
		return nil, err
	}
	return dayTemplate, err
}

func (r DayTemplateRepository) FindAllConsultsForDayTemplate(
	ctx context.Context,
	dayTemplateID uint,
) ([]*model.DayTemplateConsult, error) {
	return gorm.G[*model.DayTemplateConsult](r.db).
		Where("day_template_id = ?", dayTemplateID).
		Find(ctx)
}

func (r DayTemplateRepository) DeleteAllConsultsForDayTemplate(
	ctx context.Context,
	dayTemplateID uint,
) error {
	_, err := gorm.G[model.DayTemplateConsult](r.db).
		Where("day_template_id = ?", dayTemplateID).
		Delete(ctx)
	return err
}

func (r DayTemplateRepository) SaveDayTemplateConsults(
	ctx context.Context,
	consults []*model.DayTemplateConsult,
) ([]*model.DayTemplateConsult, error) {
	err := gorm.G[[]*model.DayTemplateConsult](r.db).Create(ctx, &consults)
	if err != nil {
		return nil, err
	}
	return consults, nil
}
