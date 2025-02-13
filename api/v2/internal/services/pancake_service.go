package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type PancakeService struct {
	Repo *repositories.PancakeRepository
}

func NewPancakeService(repo *repositories.PancakeRepository) *PancakeService {
	return &PancakeService{Repo: repo}
}

func (s *PancakeService) Create(user_id uuid.UUID, note *models.Pancake) error {
	note.UserID = user_id
	return s.Repo.Create(note)
}

func (s *PancakeService) GetByID(userID uuid.UUID, noteID int) (*models.Pancake, error) {
	note, err := s.Repo.FindByID(noteID, userID)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *PancakeService) GetAll(userID uuid.UUID) ([]models.Pancake, error) {
	return s.Repo.GetAll(userID)
}

func (s *PancakeService) Update(userID uuid.UUID, new *models.Pancake) (error, *models.Pancake) {
	note, err := s.Repo.FindByID(new.NoteID, userID)
	if err != nil {
		return err, nil
	}

	note.Title = new.Title
	note.Content = new.Content
	note.Emoji = new.Emoji
	note.Color = new.Color

	err = s.Repo.Update(note)
	return err, note
}

func (s *PancakeService) Delete(userID uuid.UUID, noteID int) error {
	exists, err := s.Repo.Exists(noteID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(noteID)
}
