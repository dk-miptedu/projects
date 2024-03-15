package handlers

import (
	"database/sql"
	"encoding/json"
	"lessons/models"
	"lessons/repo"
	"net/http"
)

func Item(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var item models.Item
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := repo.Create(item, db); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(models.ItemResponse{Item: item, Ok: true})

		case "GET":
			id := r.URL.Query().Get("id")
			if id != "" {
				item := repo.Read(id, db)
				if item != nil {
					json.NewEncoder(w).Encode(models.ItemResponse{Item: *item, Ok: true})
				} else {
					http.NotFound(w, r)
				}
			} else {
				json.NewEncoder(w).Encode(models.ItemResponse{Item: models.Item{}, Ok: false})
			}
		}
	}
}

// Пример с использованием InMemoryDB
// var InMemoryDB = repo.NewInMemoryDB()

// func Item(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		var item models.Item
// 		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 		}

// 		InMemoryDB.Create(item)

// 		json.NewEncoder(w).Encode(models.ItemResponse{Item: item, Ok: true})

// 	case "GET":
// 		id := r.URL.Query().Get("id")

// 		if id != "" {
// 			item, err := InMemoryDB.Read(id)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 			}

// 			json.NewEncoder(w).Encode(models.ItemResponse{Item: *item, Ok: true})
// 		} else {
// 			items := InMemoryDB.List()
// 			json.NewEncoder(w).Encode(models.ListResponse{Item: items, Ok: true})
// 		}
// 	}
// }
