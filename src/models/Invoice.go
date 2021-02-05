package models

import "time"

type Invoice struct {
	Id          uint      `json:"id" ;gorm:"primaryKey;autoIncrement"`
	Description string    `json:"description"`
	Customer    uint      `json:"customer"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}
