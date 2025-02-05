package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EspressoRepository struct {
	DB *gorm.DB
}

func NewEspressoRepository(db *gorm.DB) *EspressoRepository {
	return &EspressoRepository{DB: db}
}

func (r *EspressoRepository) Create(session *models.EspressoSession) error {
	return r.DB.Create(session).Error
}

func (r *EspressoRepository) FindByID(id int, userID uuid.UUID) (*models.EspressoSession, error) {
	var session models.EspressoSession
	err := r.DB.Where("session_id = ? AND user_id = ?", id, userID).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *EspressoRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.EspressoSession{}).Where("session_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *EspressoRepository) GetAll(userID uuid.UUID) ([]models.EspressoSession, error) {
	var sessions []models.EspressoSession
	err := r.DB.Where("user_id = ?", userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *EspressoRepository) Update(session *models.EspressoSession) error {
	return r.DB.Save(session).Error
}

func (r *EspressoRepository) Delete(id int) error {
	return r.DB.Delete(&models.EspressoSession{}, id).Error
}
