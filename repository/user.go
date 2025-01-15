package repository

import (
	"database/sql"
	"log"
	"time"

	"transactions-app/model"
)

type UserRepository struct {
	DB *sql.DB
}

func GetUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user model.PostUser) bool {
	stmt, err := r.DB.Prepare("INSERT INTO users (phone, name, date_added) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	currentDate := time.Now()
	_, err = stmt.Exec(&user.Phone, user.Name, currentDate)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *UserRepository) SelectUsers() []model.User {
	var result []model.User
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Phone, &user.Name, &user.DateAdded)
		if err != nil {
			log.Println(err)
		} else {
			result = append(result, user)
		}
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
	}

	return result
}

func (r *UserRepository) SelectUserByPhone(phone string) (model.User, error) {
	var user model.User
	err := r.DB.QueryRow("SELECT * FROM users WHERE phone = $1", phone).Scan(&user.Phone, &user.Name, &user.DateAdded)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}
		log.Println(err)
		return model.User{}, err
	}
	return user, nil
}
