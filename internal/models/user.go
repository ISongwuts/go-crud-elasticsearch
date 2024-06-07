package models

type User struct {
	ID 			string 	`json:"user_id"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
	Age 		int		`json:"age"`
}