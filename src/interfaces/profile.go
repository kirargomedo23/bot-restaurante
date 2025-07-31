package interfaces

type Profile struct {
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
}
