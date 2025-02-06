package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type CerealService struct {
	Repo *repositories.CerealRepository
}

func NewCerealService(repo *repositories.CerealRepository) *CerealService {
	return &CerealService{Repo: repo}
}

func (s *CerealService) Create(user_id uuid.UUID, day *models.CerealDay) error {
	day.UserID = user_id
	return s.Repo.Create(day)
}

func (s *CerealService) GetByID(userID uuid.UUID, dayID int) (*models.CerealDay, error) {
	day, err := s.Repo.FindByID(dayID, userID)
	if err != nil {
		return nil, err
	}
	return day, nil
}

func (s *CerealService) GetByDate(userID uuid.UUID, dateStr string) (*models.CerealDay, error) {
	day, err := s.Repo.FindByDate(userID, dateStr)
	if err != nil {
		return nil, err
	}
	return day, nil
}

func (s *CerealService) GetAll(userID uuid.UUID) ([]models.CerealDay, error) {
	return s.Repo.GetAll(userID)
}

func (s *CerealService) Update(userID uuid.UUID, new *models.CerealDay) (error, *models.CerealDay) {
	day, err := s.Repo.FindByID(new.DayID, userID)
	if err != nil {
		return err, nil
	}

	day.Emoji = new.Emoji

	dateStr := new.Date.Format("2006-01-02")
	existingDay, err := s.Repo.FindByDate(userID, dateStr)
	if err == nil && existingDay.DayID != new.DayID {
		return fmt.Errorf("cannot update to a date that already exists for another entry"), nil
	}

	day.Date = new.Date

	for i := range day.Activities {
		day.Activities[i].Date = new.Date
	}

	err = s.Repo.Update(day)
	return err, day
}

func (s *CerealService) Delete(userID uuid.UUID, dayID int) error {
	exists, err := s.Repo.Exists(dayID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(dayID)
}

func (s *CerealService) CreateActivity(user_id uuid.UUID, day_id int, activity *models.CerealActivity) (*models.CerealActivity, error) {
	day, err := s.Repo.FindByID(day_id, user_id)
	if err != nil {
		return nil, err
	}
	activity.DayID = day.DayID
	activity.Date = day.Date
	err = s.Repo.CreateActivity(activity)
	return activity, err
}

func (s *CerealService) GetActivity(userID uuid.UUID, dayID int, activityID int) (*models.CerealActivity, error) {
	day, err := s.Repo.FindByID(dayID, userID)
	if err != nil {
		return nil, err
	}

	for _, activity := range day.Activities {
		if activity.ActivityID == activityID {
			return &activity, nil
		}
	}

	return nil, fmt.Errorf("activity not found")
}

func (s *CerealService) UpdateActivity(userID uuid.UUID, dayID int, activityID int, newData *models.CerealActivity) (*models.CerealActivity, error) {
	day, err := s.Repo.FindByID(dayID, userID)
	if err != nil {
		return nil, err
	}

	var activity *models.CerealActivity
	for i := range day.Activities {
		if day.Activities[i].ActivityID == activityID {
			activity = &day.Activities[i]
			break
		}
	}

	if activity == nil {
		return nil, fmt.Errorf("activity not found")
	}

	activity.Title = newData.Title
	activity.Date = newData.Date
	activity.StartTime = newData.StartTime
	activity.EndTime = newData.EndTime
	activity.Notify = newData.Notify

	err = s.Repo.UpdateActivity(activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (s *CerealService) DeleteActivity(userID uuid.UUID, dayID int, activityID int) error {
	day, err := s.Repo.FindByID(dayID, userID)
	if err != nil {
		return err
	}

	return s.Repo.DeleteActivity(day.DayID, activityID)
}
