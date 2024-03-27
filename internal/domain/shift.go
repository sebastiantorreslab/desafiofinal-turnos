package domain

type Shift struct {
	ID        int    `json:"ID"`
	ShiftDate string `json:"shift_date" binding:"required"`
	ShiftHour string `json:"shift_hour" binding:"required"`
	IdPatient int    `json:"id_patient" binding:"required"`
	IdDentist int    `json:"id_dentist" binding:"required"`
}
