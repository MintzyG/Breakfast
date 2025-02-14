package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type YogurtService struct {
	Repo *repositories.YogurtRepository
}

func NewYogurtService(repo *repositories.YogurtRepository) *YogurtService {
	return &YogurtService{Repo: repo}
}

func (s *YogurtService) Create(user_id uuid.UUID, task *models.Yogurt) error {
	task.UserID = user_id
	return s.Repo.Create(task)
}

func (s *YogurtService) GetByID(userID uuid.UUID, taskID int) (*models.Yogurt, error) {
	task, err := s.Repo.FindByID(taskID, userID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *YogurtService) GetAll(userID uuid.UUID) ([]models.Yogurt, error) {
	return s.Repo.GetAll(userID)
}

func (s *YogurtService) Update(userID uuid.UUID, new *models.Yogurt) (error, *models.Yogurt) {
	task, err := s.Repo.FindByID(new.TaskID, userID)
	if err != nil {
		return err, nil
	}

	task.Title = new.Title
	task.Description = new.Description
	task.Emoji = new.Emoji
	task.Color = new.Color
	task.Priority = new.Priority
	task.TaskSize = new.TaskSize
	task.Completed = new.Completed
	task.Difficulty = new.Difficulty

	err = s.Repo.Update(task)
	return err, task
}

func (s *YogurtService) UpdateCompleted(userID uuid.UUID, new *models.Yogurt) (error, *models.Yogurt) {
	task, err := s.Repo.FindByID(new.TaskID, userID)
	if err != nil {
		return err, nil
	}

	task.Completed = new.Completed

	err = s.Repo.Update(task)
	return err, task
}

func (s *YogurtService) Delete(userID uuid.UUID, taskID int) error {
	exists, err := s.Repo.Exists(taskID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(taskID)
}
