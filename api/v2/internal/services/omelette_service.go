package services

import (
	"breakfast/internal/models"
	"breakfast/internal/repositories"
	"fmt"

	"github.com/google/uuid"
)

type OmeletteService struct {
	Repo *repositories.OmeletteRepository
}

func NewOmeletteService(repo *repositories.OmeletteRepository) *OmeletteService {
	return &OmeletteService{Repo: repo}
}

func (s *OmeletteService) Create(user_id uuid.UUID, table *models.OmeletteTable) error {
	table.UserID = user_id
	return s.Repo.Create(table)
}

func (s *OmeletteService) GetByID(userID uuid.UUID, eventID int) (*models.OmeletteTable, error) {
	table, err := s.Repo.FindByID(eventID, userID)
	if err != nil {
		return nil, err
	}
	return table, nil
}

func (s *OmeletteService) GetAll(userID uuid.UUID) ([]models.OmeletteTable, error) {
	return s.Repo.GetAll(userID)
}

func (s *OmeletteService) Update(userID uuid.UUID, new *models.OmeletteTable) (error, *models.OmeletteTable) {
	table, err := s.Repo.FindByID(new.TableID, userID)
	if err != nil {
		return err, nil
	}

	table.TableName = new.TableName
	table.Emoji = new.Emoji

	err = s.Repo.Update(table)
	return err, table
}

func (s *OmeletteService) Delete(userID uuid.UUID, table_id int) error {
	exists, err := s.Repo.Exists(table_id, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.Delete(table_id)
}

func (s *OmeletteService) CreateList(user_id uuid.UUID, table_id int, list *models.OmeletteList) (*models.OmeletteList, error) {
	table, err := s.Repo.FindByID(table_id, user_id)
	if err != nil {
		return nil, err
	}

	list.TableID = table.TableID
	list.UserID = table.UserID
  list.Position = len(table.Lists)

	err = s.Repo.CreateList(list)
	return list, err
}

func (s *OmeletteService) GetListByID(userID uuid.UUID, list_id int) (*models.OmeletteList, error) {
	list, err := s.Repo.FindListByID(list_id, userID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *OmeletteService) GetAllLists(userID uuid.UUID) ([]models.OmeletteList, error) {
	return s.Repo.GetAllLists(userID)
}

func (s *OmeletteService) UpdateList(userID uuid.UUID, table_id int, list_id int, newData *models.OmeletteList) (*models.OmeletteList, error) {
	table, err := s.Repo.FindByID(table_id, userID)
	if err != nil {
		return nil, err
	}

	var list *models.OmeletteList
	for i := range table.Lists {
		if table.Lists[i].ListID == list_id {
			list = &table.Lists[i]
			break
		}
	}

	if list == nil {
		return nil, fmt.Errorf("reminder not found")
	}

	list.ListName = newData.ListName
	list.Position = newData.Position

	err = s.Repo.UpdateList(list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *OmeletteService) DeleteList(userID uuid.UUID, table_id int, list_id int) error {
	exists, err := s.Repo.ListExists(table_id, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.DeleteList(list_id)
}

func (s *OmeletteService) CreateCard(user_id uuid.UUID, list_id int, card *models.OmeletteCard) (*models.OmeletteCard, error) {
	list, err := s.Repo.FindListByID(list_id, user_id)
	if err != nil {
		return nil, err
	}

	card.TableID = list.TableID
	card.ListID = list.ListID
	card.UserID = list.UserID
  card.Position = len(list.Cards)

	err = s.Repo.CreateCard(card)
	return card, err
}

func (s *OmeletteService) GetCardByID(userID uuid.UUID, card_id int) (*models.OmeletteCard, error) {
	card, err := s.Repo.FindCardByID(card_id, userID)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *OmeletteService) GetAllCards(userID uuid.UUID) ([]models.OmeletteCard, error) {
	return s.Repo.GetAllCards(userID)
}

func (s *OmeletteService) UpdateCard(userID uuid.UUID, card_id int, newData *models.OmeletteCard) (*models.OmeletteCard, error) {
	card, err := s.Repo.FindCardByID(card_id, userID)
	if err != nil {
		return nil, err
	}

	card.Position = newData.Position
	card.CardName = newData.CardName
	card.Content = newData.Content
	card.ListID = newData.ListID

	err = s.Repo.UpdateCard(card)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (s *OmeletteService) DeleteCard(userID uuid.UUID, card_id int) error {
	exists, err := s.Repo.CardExists(card_id, userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Model doesn't belong to user or exists")
	}
	return s.Repo.DeleteCard(card_id)
}
