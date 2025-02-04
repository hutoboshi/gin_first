package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

type ItemMemotyRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemotyRepository{items: items}
}

func (r *ItemMemotyRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemotyRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}
