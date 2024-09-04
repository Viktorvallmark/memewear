package dbstore

import (
	"container/list"
	"goth/internal/store"

	"gorm.io/gorm"
)

type CartStore struct {
	db *gorm.DB
}

type NewCartStoreParams struct {
	DB *gorm.DB
}

func NewCartStore(params NewCartStoreParams) *CartStore {
	return &CartStore{
		db: params.DB,
	}
}

func (s *CartStore) CreateCart(amount float32, user store.User) error {

	return s.db.Create(&store.Cart{
		Amount: amount,
		User:   user,
		Items:  *list.New(),
	}).Error
}

func (s *CartStore) GetCart(user store.User) (*store.Cart, error) {

	var cart store.Cart
	err := s.db.Where("user = ?", user).First(&cart).Error

	if err != nil {
		return nil, err
	}
	return &cart, err
}

func (c *CartStore) AddItemToCart(user store.User, item store.Item) (store.Item, error) {
	var cart store.Cart
	err := c.db.Where("user = ?", user).First(&cart).Error
	if err != nil {
		return err
	}
	cart.Items.PushBack(item)
	return item, nil
}
