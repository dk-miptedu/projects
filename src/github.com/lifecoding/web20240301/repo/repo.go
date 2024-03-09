package repo

import (
	"errors"
	"sync"
	"web20240301/models"
)

var (
	errNotFound = errors.New("item not found")
)

type InMemoryDB struct {
	items map[string]models.Item
	mu    sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		items: make(map[string]models.Item),
	}
}

func (db *InMemoryDB) Create(item models.Item) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.items[item.ID] = item

}

func (db *InMemoryDB) Read(id string) (*models.Item, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	item, exist := db.items[id]
	if !exist {
		return nil, errNotFound
	}
	array
	return &item, nil

}

func (db *InMemoryDB) Update(id string, newValue string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	item, exist := db.items[id]
	if !exist {
		return errNotFound
	}
	item.Value = newValue
	db.items[id] = item
	return nil

}

func (db *InMemoryDB) Delete(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, exist := db.items[id]
	if !exist {
		return errNotFound
	}
	delete(db.items, id)

	return nil

}

func (db *InMemoryDB) List() []models.Item {
	db.mu.RLock()
	defer db.mu.RUnlock()
	items := make([]models.Item, 0, len(db.items))
	for _, item := range db.items {
		items = append(items, item)
	}
	return items
}
