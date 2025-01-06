package schemas

import "time"

type CreateEventBody struct {
	Name     string    `json:"name" validate:"required"`
	Date     time.Time `json:"date" validate:"required"`
	Location string    `json:"location" validate:"required"`
}

type UpdateEventBody struct {
	Name     string    `json:"name"`
	Date     *time.Time `json:"date"`
	Location string    `json:"location"`
}