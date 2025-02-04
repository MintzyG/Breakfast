package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type YogurtRepository struct {
	DB *gorm.DB
}

func NewYogurtRepository(db *gorm.DB) *YogurtRepository {
	return &YogurtRepository{DB: db}
}

func (r *YogurtRepository) Create(note *models.Yogurt) error {
	return r.DB.Create(note).Error
}

func (r *YogurtRepository) FindByID(id int, userID uuid.UUID) (*models.Yogurt, error) {
	var yogurt models.Yogurt
	err := r.DB.Where("task_id = ? AND user_id = ?", id, userID).First(&yogurt).Error
	if err != nil {
		return nil, err
	}
	return &yogurt, nil
}

func (r *YogurtRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Yogurt{}).Where("task_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *YogurtRepository) FindByUserID(userID uuid.UUID) ([]models.Yogurt, error) {
	var yogurts []models.Yogurt
	err := r.DB.Where("user_id = ?", userID).Find(&yogurts).Error
	if err != nil {
		return nil, err
	}
	return yogurts, nil
}

func (r *YogurtRepository) Update(task *models.Yogurt) error {
	return r.DB.Save(task).Error
}

func (r *YogurtRepository) UpdateStatus(id int, userID uuid.UUID, completed bool) error {
	return r.DB.Model(&models.Yogurt{}).
		Where("task_id = ? AND user_id = ?", id, userID).
		Update("completed", completed).Error
}

func (r *YogurtRepository) Delete(id int) error {
	return r.DB.Delete(&models.Yogurt{}, id).Error
}
