package database

type Menu struct {
	Active      bool    `json:"active"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
}
