package domain

import "time"

type Shift struct {
	ID        int       `json:"ID"`
	Patient   Patient   `json:"patient" binding:"required"`
	Dentist   Dentist   `json:"dentist" binding:"required"`
	ShiftHour time.Time `json:"shiftHour" binding:"required"`
	ShiftDate time.Time `json:"shiftDate" binding:"required"`
}
