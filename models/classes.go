package models

import "time"

type Class struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"startDate"`
	Capacity int       `json:"capacity"`
}
