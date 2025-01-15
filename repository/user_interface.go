package repository

import "transactions-app/model"

type UserRepositoryInterface interface {
	CreateUser(post model.PostUser) bool
	SelectUsers() []model.User
	SelectUserByPhone(phone string) (model.User, error)
}
