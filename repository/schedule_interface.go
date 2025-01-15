package repository

import "transactions-app/model"

type ScheduleRepositoryInterface interface {
	AddToQueue(phone string, priority int) error
	RemoveFromQueue() (*model.Schedule, error)
	GetQueue() ([]model.Schedule, error)
}
