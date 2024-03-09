package repo

import (
	"dz-05_RESTful_API/models"
	"errors"
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

func (db *InMemoryDB) CreateItem(item models.Transaction) string {
	db.mu.Lock()
	defer db.mu.Unlock()
	currentId += 1
	item.ID = strconv.Itoa(currentId)
	db.Items = append(db.Items, item)

	//fmt.Printf("item: %s\n", item)

	return strconv.Itoa(currentId)

}
func (db *InMemoryDB) UpdateItem(id string, item models.Transaction) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for index, itemBefore := range db.Items {
		//fmt.Println(index, " : ", itemBefore)
		if itemBefore.ID == id {
			item.ID = id
			db.Items = append(db.Items[:index], db.Items[index+1:]...)
			db.Items = append(db.Items, item)
			return true

		}
	}
	return false

}
func (db *InMemoryDB) DeleteItem(id string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for index, item := range db.Items {
		//fmt.Println(index, " : ", item)
		if item.ID == id {
			db.Items = append(db.Items[:index], db.Items[index+1:]...)
			return true
			break
		}

	}
	return false

}
