package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"web20240301/models"
	"web20240301/repo"
)

var InMemoryDB = repo.NewInMemoryDB()

func Item(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		var item models.Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		InMemoryDB.Create(item)

		json.NewEncoder(w).Encode(models.ItemResponse{Item: item, Ok: true})
	case "GET":
		id := r.URL.Query().Get("id") //query rq

		if id == "" {
			http.Error(w, errors.New("id can not empty").Error(), http.StatusBadRequest)
		} else if id != "" {
			item, err := InMemoryDB.Read((id))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

			}
			json.NewEncoder(w).Encode(models.ItemResponse{Item: *item, Ok: true})
		} else {
			items := InMemoryDB.List()
			json.NewEncoder(w).Encode(models.ListResponse{Item: items, Ok: true})
		}

	}
}
