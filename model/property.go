package model

import (
	"database/sql"
	"fmt"
)

// Property struct
type Property struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// GetPropertyByUserID return based on userID
func GetPropertyByUserID(db *sql.DB, userID int) (property Property) {

	query := `SELECT id, name, price FROM properties WHERE userId = ?`
	row := db.QueryRow(query, userID)
	err := row.Scan(&property.ID, &property.Name, &property.Price)

	if err != nil {
		fmt.Println(err)
	}

	return property
}
