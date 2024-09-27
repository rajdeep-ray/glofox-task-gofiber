package models

import "time"

type Booking struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Date  time.Time `json:"date"`
	Class *Class    `json:"class"`
}
