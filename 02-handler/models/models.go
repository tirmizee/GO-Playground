package models

import "time"

type ProductRequest struct {
	ProductName string `json:"productName"`
	ProductType string `json:"productType"`
}

type ProductResponse struct {
	ProductId   int       `json:"productId"`
	ProductName string    `json:"productName"`
	ProductType string    `json:"productType"`
	CreateDate  time.Time `json:"createDate"`
}

type CustomerRequest struct {
	Ids []int `json:"ids"`
}

type CustomerResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Status     bool      `json:"status"`
	CreateDate time.Time `json:"createDate"`
}

type ResponseError struct {
	Message string `json:"productId"`
}

func (r *ResponseError) Default() *ResponseError {
	if r.Message == "" {
		r.Message = "Error"
	}
	return r
}
