package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MapleRepository struct {
	DB *gorm.DB
}

func NewMapleRepository(db *gorm.DB) *MapleRepository {
	return &MapleRepository{DB: db}
}

func (r *MapleRepository) Create(habit *models.Maple) error {
	return r.DB.Create(habit).Error
}

func (r *MapleRepository) CreateDay(day *models.MapleDay) error {
	return r.DB.Create(day).Error
}

func (r *MapleRepository) FindByID(id int, userID uuid.UUID) (*models.Maple, error) {
	var maple models.Maple
	err := r.DB.Preload("MapleDays").Where("habit_id = ? AND user_id = ?", id, userID).First(&maple).Error
	if err != nil {
		return nil, err
	}
	return &maple, nil
}

func (r *MapleRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.Maple{}).Where("habit_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *MapleRepository) FindByUserID(userID uuid.UUID) ([]models.Maple, error) {
	var maples []models.Maple
	err := r.DB.Preload("MapleDays").Where("user_id = ?", userID).Find(&maples).Error
	if err != nil {
		return nil, err
	}
	return maples, nil
}

func (r *MapleRepository) Update(habit *models.Maple) error {
	return r.DB.Omit("MapleDays").Save(habit).Error
}

func (r *MapleRepository) Delete(id int) error {
	return r.DB.Delete(&models.Maple{}, id).Error
}
