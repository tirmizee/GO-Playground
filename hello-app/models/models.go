package models

import "time"

type Customer struct {
	Id         int32
	FirstName  string
	LastName   string
	CreateDate time.Time
}

type Employee struct {
	EmpId       int32     `json:"empId"`
	EmpFullname string    `json:"empFullname"`
	IsActive    bool      `json:"isActive"`
	CreateDate  time.Time `json:"createDate"`
}
