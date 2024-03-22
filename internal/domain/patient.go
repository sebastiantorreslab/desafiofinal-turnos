package domain

import "time"

type Patient struct {
	ID            int       `json:"ID"`
	DNI           int       `json:"DNI" binding:"required"`
	Name          string    `json:"name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Address       string    `json:"address" binding:"required"`
	AdmissionHour time.Time `json:"admission_hour" binding:"required"`
}
