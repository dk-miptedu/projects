package repo

import (
	"database/sql"
	"fmt"
	"golang_hws/db"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Создаём таблицу users
	createExtensionsChkPass := `
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
	`
	_, err = db.Exec(createExtensionsChkPass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Extension 'chkpass' created successfully!")

	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users
		(
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL,
		u_password text NOT NULL, 
		unique(email)
		)
	`
	createTransactionsTableQuery := `
		CREATE TABLE IF NOT EXISTS transactions
		(
		id SERIAL PRIMARY KEY,
		u_id int NOT NULL,
		amount NUMERIC NOT NULL, 
		t_currency varchar(3), 
		t_type varchar(15), 
		t_category varchar(50), 
		t_date varchar(15), 
		description VARCHAR(100) NOT NULL,
		commission NUMERIC default 0,
		CONSTRAINT transactions_u_id FOREIGN KEY (u_id) REFERENCES "users"(u_id) ON DELETE RESTRICT ON UPDATE CASCADE
	)`
	_, err = db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully!")
	_, err = db.Exec(createTransactionsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'transactions' created successfully!")
}
