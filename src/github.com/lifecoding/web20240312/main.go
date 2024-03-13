package main

//Пример использования драйвера
//Рассмотрим работу с драйвером PostgreSQL:
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Пример импорта драйвера PostgreSQL
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mplspsql"
	password = "mplspsql"
	dbname   = "mplspsql"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Создаём таблицу users
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users
		(
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT
		NULL,
		email VARCHAR(100) NOT NULL
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully!")
	result, err := db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", "john_doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted %d rows into the 'users' table.\n", rowsAffected)
}
