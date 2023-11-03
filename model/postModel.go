package model

type Book struct {
	ID       string `json:"id" example:"1"`
	Title    string `json:"title" example:"Feb's Tutor"`
	Author   string `json:"author" example:"Febs"`
	Quantity int    `json:"quantity" example:"10"`
}
