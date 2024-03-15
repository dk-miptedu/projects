package repo

import (
	"database/sql"
	"fmt"
	configs "lessons/config"
	"lessons/models"

	_ "github.com/lib/pq"
)

func InitDB(config *configs.Config) (*sql.DB, error) {
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		fmt.Println("Open err")
		return nil, err
	}

	createDB := `DROP TABLE IF EXISTS items;
	CREATE TABLE IF NOT EXISTS items (
		item_id VARCHAR(255),
		value VARCHAR(255)
	);
	`

	_, err = db.Exec(createDB)
	if err != nil {
		fmt.Println("Exec err")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping err")
		return nil, err
	}

	return db, nil
}

func Create(item models.Item, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO items (item_id, value) VALUES ($1, $2)", item.ID, item.Value)
	if err != nil {
		fmt.Println("Exec INSERT")
		return err
	}

	return nil
}

func Read(id string, db *sql.DB) *models.Item {
	var result models.Item
	db.QueryRow("SELECT item_id, value FROM items WHERE item_id = $1", id).Scan(&result)

	return &result
}

// Пример с использованием InMemoryDB
// var (
// 	errNotFound = errors.New("item not found")
// )

// type InMemoryDB struct {
// 	items map[string]models.Item
// 	mu    sync.RWMutex
// }

// func NewInMemoryDB() *InMemoryDB {
// 	return &InMemoryDB{
// 		items: make(map[string]models.Item),
// 	}
// }

// func (db *InMemoryDB) Create(item models.Item) {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()

// 	db.items[item.ID] = item
// }

// func (db *InMemoryDB) Read(id string) (*models.Item, error) {
// 	db.mu.RLock()
// 	defer db.mu.RUnlock()

// 	item, exist := db.items[id]
// 	if !exist {
// 		return nil, errNotFound
// 	}

// 	return &item, nil
// }

// func (db *InMemoryDB) Update(id string, newValue string) error {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()

// 	item, exist := db.items[id]
// 	if !exist {
// 		return errNotFound
// 	}

// 	item.Value = newValue
// 	db.items[id] = item

// 	return nil
// }

// func (db *InMemoryDB) Delete(id string) error {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()

// 	_, exist := db.items[id]
// 	if !exist {
// 		return errNotFound
// 	}

// 	delete(db.items, id)

// 	return nil
// }

// func (db *InMemoryDB) List() []models.Item {
// 	db.mu.RLock()
// 	defer db.mu.RUnlock()

// 	items := make([]models.Item, 0, len(db.items))
// 	for _, item := range db.items {
// 		items = append(items, item)
// 	}

// 	return items
// }
