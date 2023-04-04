package modelo

import "time"

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding: "required"`
	Quantity    int       `json:"quantity" binding: "required"`
	CodeValue   string    `json:"codevalue" binding: "required"`
	IsPublished bool      `json:"ispublished"`
	Expiration  time.Time `json:"expiration" binding: "required" time_format: "2006-01-02"`
	Price       float64   `json:"price" binding: "required"`
}
