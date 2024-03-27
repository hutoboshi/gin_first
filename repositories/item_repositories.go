package repositories

import "gin-fleamarket/models"

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
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
