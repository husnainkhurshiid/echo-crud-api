package model

type Employee struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Department string `json:"department"`
	Email      string `json:"email"`
}
