package models

type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type ItemResponse struct {
	Item Item `json:"item"`
	Ok   bool `json:"ok"`
}

type ListResponse struct {
	Item []Item `json:"items"`
	Ok   bool   `json:"ok"`
}
