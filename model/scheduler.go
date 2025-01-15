package model

import "time"

type Schedule struct {
	Id            uint      `json:"id"`
	Phone         string    `json:"phone"`
	Priority      int       `json:"priority"`
	ScheduledTime time.Time `json:"scheduled_time"`
}

type PostSchedule struct {
	Phone    string `json:"phone" binding:"required,e164"`
	Priority int    `json:"priority" binding:"required"`
}
