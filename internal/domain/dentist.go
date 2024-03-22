package domain

type Dentist struct {
	ID       int    `json:"ID" binding:"required"`
	License  string `json:"license" binding:"required"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
}
