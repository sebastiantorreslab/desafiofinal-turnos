package domain

type Patient struct {
	ID            int    `json:"ID"`
	DNI           int    `json:"DNI" binding:"required"`
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	AdmissionDate string `json:"admission_date" binding:"required"`
}
