package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type ToastService struct {
	Repo *repositories.ToastRepository
}

func NewToastService(repo *repositories.ToastRepository) *ToastService {
	return &ToastService{Repo: repo}
}

func (s *ToastService) Create(user_id uuid.UUID, session *models.ToastSession) error {
	session.UserID = user_id
	if session.EndTime.Before(session.StartTime) {
		return fmt.Errorf("End-time cannot be before Start-time")
	}

	session.Duration = int64(session.EndTime.Sub(session.StartTime).Seconds())
	return s.Repo.Create(session)
}

func (s *ToastService) GetByID(userID uuid.UUID, sessionID int) (*models.ToastSession, error) {
	session, err := s.Repo.FindByID(sessionID, userID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *ToastService) GetAll(userID uuid.UUID) ([]models.ToastSession, error) {
	return s.Repo.GetAll(userID)
}

func (s *ToastService) Update(userID uuid.UUID, new *models.ToastSession) (error, *models.ToastSession) {
	session, err := s.Repo.FindByID(new.SessionID, userID)
	if err != nil {
		return err, nil
	}

	session.Emoji = new.Emoji
	session.SessionName = new.SessionName
	session.Description = new.Description
	session.StartTime = new.StartTime
	session.EndTime = new.EndTime

	if session.EndTime.Before(session.StartTime) {
		return fmt.Errorf("End-time cannot be before Start-time"), session
	}

	session.Duration = int64(session.EndTime.Sub(session.StartTime).Seconds())

	err = s.Repo.Update(session)
	return err, session
}

func (s *ToastService) Delete(userID uuid.UUID, sessionID int) error {
	exists, err := s.Repo.Exists(sessionID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(sessionID)
}
