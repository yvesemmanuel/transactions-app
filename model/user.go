package model

import "time"

type User struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	DateAdded time.Time `json:"date_added"`
}

type PostUser struct {
	Phone string `json:"phone" binding:"required,e164"`
	Name  string `json:"name" binding:"required"`
}
