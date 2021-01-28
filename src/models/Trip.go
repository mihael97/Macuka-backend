package models

type Trip struct {
	Id    uint   `json:"id",gorm:"primaryKey, autoIncrement"`
	Path  string `json:"path"`
	Start uint   `json:"start"`
	End   uint   `json:"end"`
	Car   uint
}
