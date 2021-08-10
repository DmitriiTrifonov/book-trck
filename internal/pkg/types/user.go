package types

import "github.com/DmitriiTrifonov/book-trck/internal/pkg/models"

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *User) MapToDTO() *models.User {
	return &models.User{
		Username: u.UserName,
		Password: u.Password,
	}
}
