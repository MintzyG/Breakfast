package repositories

import (
	"breakfast/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OmeletteRepository struct {
	DB *gorm.DB
}

func NewOmeletteRepository(db *gorm.DB) *OmeletteRepository {
	return &OmeletteRepository{DB: db}
}

func (r *OmeletteRepository) Create(event *models.OmeletteTable) error {
	return r.DB.Create(event).Error
}

func (r *OmeletteRepository) FindByID(id int, userID uuid.UUID) (*models.OmeletteTable, error) {
	var table models.OmeletteTable
	err := r.DB.Preload("Lists.Cards").Where("table_id = ? AND user_id = ?", id, userID).First(&table).Error
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (r *OmeletteRepository) Exists(id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.OmeletteTable{}).Where("table_id = ? AND user_id = ?", id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *OmeletteRepository) GetAll(userID uuid.UUID) ([]models.OmeletteTable, error) {
	var tables []models.OmeletteTable
	err := r.DB.Where("user_id = ?", userID).Find(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *OmeletteRepository) Update(event *models.OmeletteTable) error {
	return r.DB.Omit("Lists").Save(event).Error
}

func (r *OmeletteRepository) Delete(id int) error {
	return r.DB.Delete(&models.OmeletteTable{}, id).Error
}

func (r *OmeletteRepository) CreateList(event *models.OmeletteList) error {
	return r.DB.Create(event).Error
}

func (r *OmeletteRepository) FindListByID(list_id int, userID uuid.UUID) (*models.OmeletteList, error) {
	var list models.OmeletteList
	err := r.DB.Preload("Cards").Where("list_id = ? AND user_id = ?", list_id, userID).First(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *OmeletteRepository) ListExists(list_id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.OmeletteList{}).Where("list_id = ? AND user_id = ?", list_id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *OmeletteRepository) GetAllLists(userID uuid.UUID) ([]models.OmeletteList, error) {
	var lists []models.OmeletteList
	err := r.DB.Where("user_id = ?", userID).Find(&lists).Error
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *OmeletteRepository) UpdateList(list *models.OmeletteList) error {
	return r.DB.Omit("Cards").Save(list).Error
}

func (r *OmeletteRepository) DeleteList(list_id int) error {
	return r.DB.Delete(&models.OmeletteList{}, list_id).Error
}

func (r *OmeletteRepository) CreateCard(event *models.OmeletteCard) error {
	return r.DB.Create(event).Error
}

func (r *OmeletteRepository) FindCardByID(card_id int, userID uuid.UUID) (*models.OmeletteCard, error) {
	var card models.OmeletteCard
	err := r.DB.Where("card_id = ? AND user_id = ?", card_id, userID).First(&card).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *OmeletteRepository) CardExists(card_id int, userID uuid.UUID) (bool, error) {
	var count int64
	err := r.DB.Model(&models.OmeletteCard{}).Where("card_id = ? AND user_id = ?", card_id, userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *OmeletteRepository) GetAllCards(userID uuid.UUID) ([]models.OmeletteCard, error) {
	var cards []models.OmeletteCard
	err := r.DB.Where("user_id = ?", userID).Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (r *OmeletteRepository) UpdateCard(card *models.OmeletteCard) error {
	return r.DB.Save(card).Error
}

func (r *OmeletteRepository) DeleteCard(card_id int) error {
	return r.DB.Delete(&models.OmeletteCard{}, card_id).Error
}
