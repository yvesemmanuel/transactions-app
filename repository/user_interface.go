package repository

import "fintech-app/model"

type UserRepositoryInterface interface {
	CreateUser(post model.PostUser) bool
	SelectUsers() []model.User
	SelectUserByID(id string) (model.User, error)
}
