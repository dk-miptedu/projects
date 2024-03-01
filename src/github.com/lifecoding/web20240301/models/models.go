package models

type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type CreateResponse struct {
	Item Item `json:"item"`
	Ok   bool `json:"Ok"`
}
