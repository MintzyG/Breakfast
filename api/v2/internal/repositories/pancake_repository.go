package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PancakeRepository struct {
	DB *gorm.DB
}

func NewPancakeRepository(db *gorm.DB) *PancakeRepository {
	return &PancakeRepository{DB: db}
}

func (r *PancakeRepository) Create(note *models.Pancake) error {
	return r.DB.Create(note).Error
}

func (r *PancakeRepository) FindByID(id int, userID uuid.UUID) (*models.Pancake, error) {
	var pancake models.Pancake
	err := r.DB.Where("note_id = ? AND user_id = ?", id, userID).First(&pancake).Error
	if err != nil {
		return nil, err
	}
	return &pancake, nil
}

func (r *PancakeRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Pancake{}).Where("note_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *PancakeRepository) GetAll(userID uuid.UUID) ([]models.Pancake, error) {
	var pancakes []models.Pancake
	err := r.DB.Where("user_id = ?", userID).Find(&pancakes).Error
	if err != nil {
		return nil, err
	}
	return pancakes, nil
}

func (r *PancakeRepository) Update(note *models.Pancake) error {
	return r.DB.Save(note).Error
}

func (r *PancakeRepository) Delete(id int) error {
	return r.DB.Delete(&models.Pancake{}, id).Error
}
