package models

type Category struct {
	ID          int    `json:"id" example:"1"`
	Name        string `json:"name" example:"Minuman"`
	Description string `json:"description" example:"Aneka Jus dan Kopi"`
}
