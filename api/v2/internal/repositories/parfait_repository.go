package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ParfaitRepository struct {
	DB *gorm.DB
}

func NewParfaitRepository(db *gorm.DB) *ParfaitRepository {
	return &ParfaitRepository{DB: db}
}

func (r *ParfaitRepository) Create(event *models.ParfaitEvent) error {
	return r.DB.Create(event).Error
}

func (r *ParfaitRepository) FindByID(id int, userID uuid.UUID) (*models.ParfaitEvent, error) {
	var event models.ParfaitEvent
	err := r.DB.Preload("Reminders").Where("event_id = ? AND user_id = ?", id, userID).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *ParfaitRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.ParfaitEvent{}).Where("event_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *ParfaitRepository) GetAll(userID uuid.UUID) ([]models.ParfaitEvent, error) {
	var events []models.ParfaitEvent
	err := r.DB.Preload("Reminders").Where("user_id = ?", userID).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *ParfaitRepository) Update(event *models.ParfaitEvent) error {
	return r.DB.Save(event).Error
}

func (r *ParfaitRepository) Delete(id int) error {
	return r.DB.Delete(&models.ParfaitEvent{}, id).Error
}

func (r *ParfaitRepository) CreateReminder(reminder *models.ParfaitReminder) error {
	return r.DB.Create(reminder).Error
}

func (r *ParfaitRepository) GetAllReminders(userID uuid.UUID) ([]models.ParfaitReminder, error) {
	var reminders []models.ParfaitReminder
	err := r.DB.Where("user_id = ?", userID).Find(&reminders).Error
	if err != nil {
		return nil, err
	}
	return reminders, nil
}

func (r *ParfaitRepository) UpdateReminder(reminder *models.ParfaitReminder) error {
	return r.DB.Save(reminder).Error
}

func (r *ParfaitRepository) DeleteReminder(event_id int, reminder_id int) error {
	return r.DB.Where("event_id = ? AND reminder_id = ?", event_id, reminder_id).Delete(&models.ParfaitReminder{}).Error
}
