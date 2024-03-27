package domain

type Shift struct {
	ID        int    `json:"ID"`
	IdPatient int    `json:"id_patient" binding:"required"`
	IdDentist int    `json:"id_dentist" binding:"required"`
	ShiftHour string `json:"shift_hour" binding:"required"`
	ShiftDate string `json:"shift_date" binding:"required"`
}
