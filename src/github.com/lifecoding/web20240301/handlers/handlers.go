package handlers

import (
	"encoding/json"
	"net/http"
	"web20240301/models"
)

var InMemoryDB = make(map[string]models.Item)

//нужно использовать Mutex

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	InMemoryDB[item.ID] = item
	json.NewEncoder(w).Encode(models.CreateResponse{Item: item, Ok: true})

}
