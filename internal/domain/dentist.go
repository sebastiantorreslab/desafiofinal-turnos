package domain

type Dentist struct {
	ID       int    `json:"ID" `
	License  string `json:"license" binding:"required"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}
