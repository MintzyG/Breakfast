package services

import (
	"breakfast/internal/repositories"

	"github.com/google/uuid"
)

type DataService struct {
	UserRepo *repositories.UserRepository
}

func NewDataService(repo *repositories.UserRepository) *DataService {
	return &DataService{UserRepo: repo}
}

func (s *DataService) Me(id uuid.UUID) (string, error) {
	user, err := s.UserRepo.FindByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
