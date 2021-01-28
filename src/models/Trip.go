package models

import "time"

type Trip struct {
	Id    uint      `json:"id",gorm:"primaryKey, autoIncrement"`
	Date  time.Time `json:"date"`
	Path  string    `json:"path"`
	Start uint      `json:"start"`
	End   uint      `json:"end"`
	Car   uint      `json:"car"`
}
