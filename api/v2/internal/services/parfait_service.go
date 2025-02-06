package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type ParfaitService struct {
	Repo *repositories.ParfaitRepository
}

func NewParfaitService(repo *repositories.ParfaitRepository) *ParfaitService {
	return &ParfaitService{Repo: repo}
}

func (s *ParfaitService) Create(user_id uuid.UUID, event *models.ParfaitEvent) error {
	event.UserID = user_id
	return s.Repo.Create(event)
}

func (s *ParfaitService) GetByID(userID uuid.UUID, eventID int) (*models.ParfaitEvent, error) {
	event, err := s.Repo.FindByID(eventID, userID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *ParfaitService) GetAll(userID uuid.UUID) ([]models.ParfaitEvent, error) {
	return s.Repo.GetAll(userID)
}

func (s *ParfaitService) Update(userID uuid.UUID, new *models.ParfaitEvent) (error, *models.ParfaitEvent) {
	event, err := s.Repo.FindByID(new.EventID, userID)
	if err != nil {
		return err, nil
	}

	event.Title = new.Title
	event.Description = new.Description
	event.StartTime = new.StartTime
	event.EndTime = new.EndTime
	event.Location = new.Location

	err = s.Repo.Update(event)
	return err, event
}

func (s *ParfaitService) Delete(userID uuid.UUID, eventID int) error {
	exists, err := s.Repo.Exists(eventID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(eventID)
}

func (s *ParfaitService) CreateReminder(user_id uuid.UUID, event_id int, reminder *models.ParfaitReminder) (*models.ParfaitReminder, error) {
	event, err := s.Repo.FindByID(event_id, user_id)
	if err != nil {
		return nil, err
	}

	reminder.EventID = event.EventID
	reminder.UserID = event.UserID

	err = s.Repo.CreateReminder(reminder)
	return reminder, err
}

func (s *ParfaitService) GetReminder(userID uuid.UUID, eventID int, reminderID int) (*models.ParfaitReminder, error) {
	event, err := s.Repo.FindByID(eventID, userID)
	if err != nil {
		return nil, err
	}

	for _, reminder := range event.Reminders {
		if reminder.ReminderID == reminderID {
			return &reminder, nil
		}
	}

	return nil, fmt.Errorf("Reminder not found")
}

func (s *ParfaitService) GetAllReminders(userID uuid.UUID) ([]models.ParfaitReminder, error) {
	return s.Repo.GetAllReminders(userID)
}

func (s *ParfaitService) UpdateReminder(userID uuid.UUID, eventID int, reminderID int, newData *models.ParfaitReminder) (*models.ParfaitReminder, error) {
	event, err := s.Repo.FindByID(eventID, userID)
	if err != nil {
		return nil, err
	}

	var reminder *models.ParfaitReminder
	for i := range event.Reminders {
		if event.Reminders[i].ReminderID == reminderID {
			reminder = &event.Reminders[i]
			break
		}
	}

	if reminder == nil {
		return nil, fmt.Errorf("reminder not found")
	}

	reminder.RemindAt = newData.RemindAt
	reminder.Title = newData.Title
	reminder.Description = newData.Description

	err = s.Repo.UpdateReminder(reminder)
	if err != nil {
		return nil, err
	}

	return reminder, nil
}

func (s *ParfaitService) DeleteReminder(userID uuid.UUID, eventID int, reminderID int) error {
	event, err := s.Repo.FindByID(eventID, userID)
	if err != nil {
		return err
	}

	return s.Repo.DeleteReminder(event.EventID, reminderID)
}
