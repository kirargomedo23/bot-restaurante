package interfaces

type Address struct {
	Active   bool   `json:"active"`
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	Province string `json:"province"`
	Region   string `json:"region"`
}
