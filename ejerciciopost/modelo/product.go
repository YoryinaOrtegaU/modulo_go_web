package modelo

import (
	"strings"
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding: "required"`
	Quantity    int       `json:"quantity" binding: "required"`
	CodeValue   string    `json:"code_value" binding: "required"`
	IsPublished bool      `json:"ispublished"`
	Expiration  CivilTime `json:"expiration" binding: "required"`
	Price       float64   `json:"price" binding: "required"`
}

type CivilTime time.Time

func (c *CivilTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("02/01/2006", value) //parse time
	if err != nil {
		return err
	}
	*c = CivilTime(t) //set result using the pointer
	return nil
}

func (c CivilTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(c).Format("02/01/2006") + `"`), nil
}
