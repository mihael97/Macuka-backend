package dto

type CustomerPairDto struct {
	Id   uint   `json:"id"`
	Iban string `json:"iban"`
	Name string `json:"name"`
}
