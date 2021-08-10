package user

import (
	"github.com/DmitriiTrifonov/book-trck/internal/pkg/database"
	"github.com/DmitriiTrifonov/book-trck/internal/pkg/helper"
	"github.com/DmitriiTrifonov/book-trck/internal/pkg/models"
	"github.com/DmitriiTrifonov/book-trck/internal/pkg/types"
	"github.com/pkg/errors"
)

type UserRepo struct {
	ds *database.DataStorage
	helper.HashHelper
}

func NewUserRepo(ds *database.DataStorage) (*UserRepo, error) {
	return &UserRepo{ds: ds}, nil
}

// Add is used to add a new user
func (ur *UserRepo) Add(user types.User) error {
	_, err := ur.Get(user)
	if err == nil {
		return errors.Errorf("user %s has already registered", user.UserName)
	}

	q := `INSERT INTO public.users (username, password)
         VALUES ($1, $2)`

	tr, err := ur.ds.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(q, user.UserName, ur.Encode(user.Password))
	if err != nil {
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) Get(u types.User) (*models.User, error) {
	q := `SELECT id, username, password FROM public.users WHERE username = $1`

	user := &models.User{}
	err := ur.ds.DB.Get(user, q, u.UserName)
	if err != nil {
		return nil, errors.Wrap(err, "user was not found")
	}
	if user.Password != ur.Encode(u.Password) {
		return nil, errors.New("password is incorrect")
	}
	return user, nil
}
