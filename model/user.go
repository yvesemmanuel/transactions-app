package model

import "time"

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	DateAdded time.Time `json:"date_added"`
}

type PostUser struct {
	Name   string  `json:"name" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
