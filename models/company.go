package models

type Company struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	Location string `json:"location"`
}
