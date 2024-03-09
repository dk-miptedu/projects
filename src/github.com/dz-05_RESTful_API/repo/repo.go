package repo

import (
	"dz-05_RESTful_API/models"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

var (
	errNotFound = errors.New("item not found")
	currentId   int
)

type InMemoryDB struct {
	Items []models.Transaction
	mu    sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Items: make([]models.Transaction, 0),
	}
}

func (db *InMemoryDB) Create(item models.Transaction) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	currentId += 1
	item.ID = strconv.Itoa(currentId)
	db.Items = append(db.Items, item)

	fmt.Printf("item: %s\n", item)

	return strconv.Itoa(currentId)

}
