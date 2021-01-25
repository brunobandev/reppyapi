package bill

import "time"

// Bill struct
type Bill struct {
	ID          int
	PropertyID  int
	Description string
	Price       float64
	File        string
	DueAt       time.Time
}
