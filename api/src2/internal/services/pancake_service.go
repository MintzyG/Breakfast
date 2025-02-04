package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"

	"github.com/google/uuid"
)

type PancakeService struct {
	Repo *repositories.PancakeRepository
}

func NewPancakeService(repo *repositories.PancakeRepository) *PancakeService {
	return &PancakeService{Repo: repo}
}

func (s *PancakeService) Create(user_id uuid.UUID, note models.Pancake) error {
	note.UserID = user_id
	return s.Repo.Create(&note)
}

func (s *PancakeService) GetNoteByID(userID uuid.UUID, noteID int) (*models.Pancake, error) {
	note, err := s.Repo.FindByID(noteID, userID)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *PancakeService) GetUserNotes(userID uuid.UUID) ([]models.Pancake, error) {
	return s.Repo.FindByUserID(userID)
}

func (s *PancakeService) UpdateNote(userID uuid.UUID, new models.Pancake) error {
	note, err := s.Repo.FindByID(new.NoteID, userID)
	if err != nil {
		return err
	}

	note.Title = new.Title
	note.Content = new.Content
	note.Emoji = new.Emoji

	return s.Repo.Update(note)
}

func (s *PancakeService) DeleteNote(userID uuid.UUID, noteID int) error {
	_, err := s.Repo.FindByID(noteID, userID)
	if err != nil {
		return err
	}
	return s.Repo.Delete(noteID)
}
