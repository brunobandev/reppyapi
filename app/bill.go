package app

import "time"

type Bill struct {
	ID int
	Property
	Description string
	Price       float64
	File        string
	DueAt       time.Time
}
