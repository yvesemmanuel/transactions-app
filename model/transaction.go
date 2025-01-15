package model

import "time"

type Transaction struct {
	Id            uint      `json:"id"`
	FromUserId    uint      `json:"from_user_id"`
	ToUserId      uint      `json:"to_user_id"`
	Amount        float64   `json:"amount"`
	DateInitiated time.Time `json:"date_initiated"`
	Status        string    `json:"status"`
}

type PostTransaction struct {
	FromUserId uint    `json:"from_user_id" binding:"required"`
	ToUserId   uint    `json:"to_user_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}
