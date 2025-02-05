package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ToastRepository struct {
	DB *gorm.DB
}

func NewToastRepository(db *gorm.DB) *ToastRepository {
	return &ToastRepository{DB: db}
}

func (r *ToastRepository) Create(session *models.ToastSession) error {
	return r.DB.Create(session).Error
}

func (r *ToastRepository) FindByID(id int, userID uuid.UUID) (*models.ToastSession, error) {
	var session models.ToastSession
	err := r.DB.Where("session_id = ? AND user_id = ?", id, userID).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *ToastRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.ToastSession{}).Where("session_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *ToastRepository) GetAll(userID uuid.UUID) ([]models.ToastSession, error) {
	var sessions []models.ToastSession
	err := r.DB.Where("user_id = ?", userID).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *ToastRepository) Update(session *models.ToastSession) error {
	return r.DB.Save(session).Error
}

func (r *ToastRepository) Delete(id int) error {
	return r.DB.Delete(&models.ToastSession{}, id).Error
}
