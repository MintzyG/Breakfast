package repositories

import (
	"breakfast/internal/models"

	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CerealRepository struct {
	DB *gorm.DB
}

func NewCerealRepository(db *gorm.DB) *CerealRepository {
	return &CerealRepository{DB: db}
}

func (r *CerealRepository) Create(day *models.CerealDay) error {
	return r.DB.Create(day).Error
}

func (r *CerealRepository) CreateActivity(day *models.CerealActivity) error {
	return r.DB.Create(day).Error
}

func (r *CerealRepository) FindByID(id int, userID uuid.UUID) (*models.CerealDay, error) {
	var day models.CerealDay
	err := r.DB.Preload("Activities").Where("day_id = ? AND user_id = ?", id, userID).First(&day).Error
	if err != nil {
		return nil, err
	}
	return &day, nil
}

func (r *CerealRepository) FindByDate(userID uuid.UUID, dateStr string) (*models.CerealDay, error) {
	var day models.CerealDay

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	err = r.DB.Preload("Activities").Where("DATE(date) = ? AND user_id = ?", date.Format("2006-01-02"), userID).First(&day).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("No routine created on this day")
		}
		return nil, err
	}

	return &day, nil
}

func (r *CerealRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.CerealDay{}).Where("day_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CerealRepository) GetAll(userID uuid.UUID) ([]models.CerealDay, error) {
	var days []models.CerealDay
	err := r.DB.Preload("Activities").Where("user_id = ?", userID).Find(&days).Error
	if err != nil {
		return nil, err
	}
	return days, nil
}

func (r *CerealRepository) Update(day *models.CerealDay) error {
	return r.DB.Save(day).Error
}

func (r *CerealRepository) Delete(id int) error {
	return r.DB.Delete(&models.CerealDay{}, id).Error
}
