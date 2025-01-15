package repository

import "transactions-app/model"

type UserRepositoryInterface interface {
	CreateUser(user model.PostUser) bool
	SelectUsers() []model.User
	SelectUserByPhone(phone string) (model.User, error)
}
