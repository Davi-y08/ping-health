package user

import "github.com/google/uuid"

type User struct {
	ID				uuid.UUID	`json:"id"`
	Name 			string 		`json:"name"`
	Email 			string 		`json:"email" gorm:"uniqueIndex;size:180"`
	PasswordHash 	string 		`json:"-"`
	Role 			string 		`json:"role" gorm:"default:user"`
}