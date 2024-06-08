package models

type Book struct {
	ID 			string	`json:"book_id"`
	Title 		string 	`json:"book_title"`
	Genre		string  `json:"book_genre"`
	Description string 	`json:"book_desc"`
	Author 		string 	`json:"book_author"`
	Price		string	`json:"book_price"`
}