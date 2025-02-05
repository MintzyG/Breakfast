package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type EspressoService struct {
	Repo *repositories.EspressoRepository
}

func NewEspressoService(repo *repositories.EspressoRepository) *EspressoService {
	return &EspressoService{Repo: repo}
}

func (s *EspressoService) Create(user_id uuid.UUID, session *models.EspressoSession) error {
	session.UserID = user_id
	if session.FocusEnd.Before(session.FocusStart) {
		return fmt.Errorf("End-time cannot be before Start-time")
	}

	session.Duration = int64(session.FocusEnd.Sub(session.FocusStart).Seconds())
	return s.Repo.Create(session)
}

// func (s *EspressoService) GetUserSettings(userID uuid.UUID) (*models.EspressoUserSettings, error) {
// 	settings, err := s.Repo.FindSettings(userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return settings, nil
// }

func (s *EspressoService) GetByID(userID uuid.UUID, sessionID int) (*models.EspressoSession, error) {
	session, err := s.Repo.FindByID(sessionID, userID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *EspressoService) GetAll(userID uuid.UUID) ([]models.EspressoSession, error) {
	return s.Repo.GetAll(userID)
}

func (s *EspressoService) Update(userID uuid.UUID, new *models.EspressoSession) (error, *models.EspressoSession) {
	session, err := s.Repo.FindByID(new.SessionID, userID)
	if err != nil {
		return err, nil
	}

	session.SessionName = new.SessionName
	session.Emoji = new.Emoji
	session.FocusStart = new.FocusStart
	session.FocusEnd = new.FocusEnd
	session.BreakTime = new.BreakTime
	session.BigBreak = new.BigBreak
	session.Rounds = new.Rounds
	session.Laps = new.Laps

	if session.FocusEnd.Before(session.FocusStart) {
		return fmt.Errorf("End-time cannot be before Start-time"), session
	}

	session.Duration = int64(session.FocusEnd.Sub(session.FocusStart).Seconds())

	err = s.Repo.Update(session)
	return err, session
}

func (s *EspressoService) Delete(userID uuid.UUID, sessionID int) error {
	exists, err := s.Repo.Exists(sessionID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(sessionID)
}
