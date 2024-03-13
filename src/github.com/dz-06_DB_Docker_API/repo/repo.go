package repo

import (
	"database/sql"
	"dz-06_DB_Docker_API/models"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
)

var (
	errNotFound = errors.New("item not found")
	currentId   int
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mplspsql"
	password = "mplspsql"
	dbname   = "transactionsdb"
)

type InMemoryDB struct {
	Items []models.Transaction
	mu    sync.RWMutex
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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
