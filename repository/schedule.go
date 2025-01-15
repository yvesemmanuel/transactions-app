package repository

import (
	"database/sql"
	"log"

	"transactions-app/model"
)

type ScheduleRepository struct {
	DB *sql.DB
}

func InstanteScheduleRepository(db *sql.DB) ScheduleRepositoryInterface {
	return &ScheduleRepository{DB: db}
}

func (r *ScheduleRepository) AddToQueue(phone string, priority int) error {
	query := "INSERT INTO schedules (phone, priority) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, phone, priority)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (r *ScheduleRepository) RemoveFromQueue() (*model.Schedule, error) {
	var schedule model.Schedule
	query := `
        DELETE FROM schedules
        WHERE id = (
            SELECT id FROM schedules
            ORDER BY priority DESC, scheduled_time ASC
            LIMIT 1
        )
        RETURNING id, phone, priority, scheduled_time
    `
	err := r.DB.QueryRow(query).Scan(&schedule.Id, &schedule.Phone, &schedule.Priority, &schedule.ScheduledTime)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &schedule, nil
}

func (r *ScheduleRepository) GetQueue() ([]model.Schedule, error) {
	query := `
        SELECT id, phone, priority, scheduled_time
        FROM schedules
        ORDER BY priority DESC, scheduled_time ASC
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		if err := rows.Scan(&schedule.Id, &schedule.Phone, &schedule.Priority, &schedule.ScheduledTime); err != nil {
			log.Println(err)
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}
