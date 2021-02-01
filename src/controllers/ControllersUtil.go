package controllers

import "fmt"

type HttpMethod string

const (
	GetMethod    HttpMethod = "GET"
	PostMethod   HttpMethod = "POST"
	DeleteMethod HttpMethod = "DELETE"
)

type PathMethodPair struct {
	Path   string
	Method HttpMethod
}

func (e *PathMethodPair) GetMethod() string {
	return fmt.Sprint(e.Method)
}
