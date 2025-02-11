package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"

	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
)

type MapleService struct {
	Repo *repositories.MapleRepository
}

func NewMapleService(repo *repositories.MapleRepository) *MapleService {
	return &MapleService{Repo: repo}
}

func (s *MapleService) Create(user_id uuid.UUID, habit *models.Maple) error {
	habit.UserID = user_id
	return s.Repo.Create(habit)
}

func (s *MapleService) GetByID(userID uuid.UUID, habitID int) (*models.Maple, error) {
	habit, err := s.Repo.FindByID(habitID, userID)
	if err != nil {
		return nil, err
	}
	return habit, nil
}

func (s *MapleService) GetAll(userID uuid.UUID) ([]models.Maple, error) {
	return s.Repo.GetAll(userID)
}

func (s *MapleService) Update(userID uuid.UUID, new *models.Maple) (error, *models.Maple) {
	habit, err := s.Repo.FindByID(new.HabitID, userID)
	if err != nil {
		return err, nil
	}

	habit.Title = new.Title
	habit.Emoji = new.Emoji
	habit.SmallestUnit = new.SmallestUnit

	err = s.Repo.Update(habit)
	return err, habit
}

func (s *MapleService) Delete(userID uuid.UUID, habitID int) error {
	exists, err := s.Repo.Exists(habitID, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(habitID)
}

// I broke time and space, need to rework this
func (s *MapleService) CreateDay(user_id uuid.UUID, habit_id int, day *models.MapleDay) (*models.Maple, error) {
	habit, err := s.Repo.FindByID(habit_id, user_id)
	if err != nil {
		return habit, err
	}

	for _, existingDay := range habit.MapleDays {
		if existingDay.CreatedAt.Format("2006-01-02") == day.CreatedAt.Format("2006-01-02") {
			return habit, fmt.Errorf("day already exists for this habit")
		}
	}

	habit.MapleDays = append(habit.MapleDays, *day)

	if len(habit.MapleDays) == 1 {
		if day.Achieved {
			habit.CurrStreak = 1
			habit.HighestStreak = 1
		} else {
			habit.CurrStreak = 0
		}
	} else {
		today := time.Now().Format("2006-01-02")
		dayBeingAdded := day.CreatedAt.Format("2006-01-02")

		if dayBeingAdded != today {
			sort.SliceStable(habit.MapleDays, func(i, j int) bool {
				return habit.MapleDays[i].CreatedAt.Before(habit.MapleDays[j].CreatedAt)
			})

			maxStreak := 0
			currentStreak := 0

			for i := 0; i < len(habit.MapleDays); i++ {
				if !habit.MapleDays[i].Achieved {
					currentStreak = 0
					continue
				}

				if i == 0 {
					currentStreak = 1
				} else {
					prevDay := habit.MapleDays[i-1].CreatedAt.Format("2006-01-02")
					currentDay := habit.MapleDays[i].CreatedAt.Format("2006-01-02")

					prevDayTime, _ := time.Parse("2006-01-02", prevDay)
					currentDayTime, _ := time.Parse("2006-01-02", currentDay)
					daysDiff := currentDayTime.Sub(prevDayTime).Hours() / 24

					if daysDiff == 1 && habit.MapleDays[i-1].Achieved {
						currentStreak++
					} else {
						currentStreak = 1
					}
				}

				if currentStreak > maxStreak {
					maxStreak = currentStreak
				}
			}

			habit.HighestStreak = maxStreak
		}

		yesterday := day.CreatedAt.AddDate(0, 0, -1).Format("2006-01-02")
		var foundYesterday bool
		var yesterdayAchieved bool

		for _, existingDay := range habit.MapleDays {
			if existingDay.CreatedAt.Format("2006-01-02") == yesterday {
				foundYesterday = true
				yesterdayAchieved = existingDay.Achieved
				break
			}
		}

		if foundYesterday && yesterdayAchieved && day.Achieved {
			habit.CurrStreak++
			if habit.CurrStreak > habit.HighestStreak {
				habit.HighestStreak = habit.CurrStreak
			}
		} else {
			if day.Achieved {
				habit.CurrStreak = 1
			} else {
				habit.CurrStreak = 0
			}
		}
	}

	err = s.Repo.CreateDay(day)
	if err != nil {
		return habit, fmt.Errorf("Error creating day in maple")
	}
	err = s.Repo.Update(habit)
	return habit, err
}

func (s *MapleService) GetDay(userID uuid.UUID, habitID int, dayID int) (*models.MapleDay, error) {
	habit, err := s.Repo.FindByID(habitID, userID)
	if err != nil {
		return nil, err
	}

	for _, day := range habit.MapleDays {
		if day.DayID == dayID {
			return &day, nil
		}
	}

	return nil, fmt.Errorf("Maple entry not found")
}

func (s *MapleService) UpdateDay(userID uuid.UUID, habitID int, dayID int, updatedData *models.MapleDay) (*models.MapleDay, error) {
	habit, err := s.Repo.FindByID(habitID, userID)
	if err != nil {
		return nil, err
	}

	for i, day := range habit.MapleDays {
		if day.DayID == dayID {
			habit.MapleDays[i].UnitsDone = updatedData.UnitsDone
			habit.MapleDays[i].Achieved = updatedData.Achieved
			habit.MapleDays[i].Date = updatedData.Date

			if err := s.Repo.UpdateDay(&habit.MapleDays[i]); err != nil {
				return nil, err
			}
			return &habit.MapleDays[i], nil
		}
	}

	return nil, fmt.Errorf("MapleDay not found")
}

func (s *MapleService) DeleteDay(userID uuid.UUID, habitID int, dayID int) error {
	habit, err := s.Repo.FindByID(habitID, userID)
	if err != nil {
		return err
	}

	for _, day := range habit.MapleDays {
		if day.DayID == dayID {
			return s.Repo.DeleteDay(dayID)
		}
	}

	return fmt.Errorf("MapleDay not found")
}
