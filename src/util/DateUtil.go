package util

import "time"

const (
	DateFormat = "2006-01-02"
)

func ConvertDate(dateString string) (*time.Time, error) {
	date, err := time.Parse(DateFormat, dateString)
	if err != nil {
		return nil, err
	}
	return &date, nil
}
