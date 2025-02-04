package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"errors"

	"github.com/google/uuid"
)

type PancakeService struct {
	Repo *repositories.PancakeRepository
}

func NewPancakeService(repo *repositories.PancakeRepository) *PancakeService {
	return &PancakeService{Repo: repo}
}

func (s *PancakeService) CreateNote(userID uuid.UUID, title, content, emoji string) error {
	if title == "" || content == "" {
		return errors.New("title and content cannot be empty")
	}

	note := &models.Pancake{
		UserID:  userID,
		Title:   title,
		Content: content,
		Emoji:   emoji,
	}

	return s.Repo.Create(note)
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

func (s *PancakeService) UpdateNote(userID uuid.UUID, noteID int, title, content, emoji string) error {
	note, err := s.Repo.FindByID(noteID, userID)
	if err != nil {
		return err
	}

	note.Title = title
	note.Content = content
	note.Emoji = emoji

	return s.Repo.Update(note)
}

func (s *PancakeService) DeleteNote(userID uuid.UUID, noteID int) error {
	_, err := s.Repo.FindByID(noteID, userID)
	if err != nil {
		return err
	}
	return s.Repo.Delete(noteID)
}

