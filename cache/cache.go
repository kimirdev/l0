package cache

import (
	"l0/db"
	"l0/models"
	"sync"
)

type LocalCache struct {
	locker sync.RWMutex
	orders map[string]models.OrderDTO
	repo   *db.Repository
}

func NewLocalCache(repo *db.Repository) *LocalCache {
	return &LocalCache{
		orders: make(map[string]models.OrderDTO),
		repo:   repo,
	}
}

func (c *LocalCache) Initialize() error {
	ordersArray, err := c.repo.Order.GetAll()

	if err != nil {
		return err
	}

	for _, el := range ordersArray {
		c.orders[el.OrderUID] = el
	}
	return nil
}

func (c *LocalCache) GetAll() []models.OrderDTO {
	v := make([]models.OrderDTO, 0, len(c.orders))

	for _, value := range c.orders {
		v = append(v, value)
	}
	return v
}

func (c *LocalCache) Get(key string) (models.OrderDTO, bool) {
	c.locker.Lock()
	defer c.locker.Unlock()
	val, exist := c.orders[key]

	if !exist {
		return val, false
	}

	return val, true
}

func (c *LocalCache) Set(order models.OrderDTO) (success bool) {
	c.locker.Lock()
	defer c.locker.Unlock()
	_, exist := c.orders[order.OrderUID]

	if exist {
		return false
	}

	c.orders[order.OrderUID] = order
	return true
}

func (c *LocalCache) IsExist(key string) bool {
	_, exist := c.orders[key]

	return exist
}
