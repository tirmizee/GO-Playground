package main 

import (
    "fmt"
    "time"
)

type BaseEntity struct {
	id             int64
	creeateDate    time.Time
	updateDate     time.Time
}

func (b BaseEntity) getId() int64 {
	return b.id
}


type UserEntity struct {
	username       string
	password       string
	BaseEntity
}

func structInheritance() {

	baseEntity := BaseEntity{
		id: 1,
		creeateDate: time.Now(),
	}

	userEntity := UserEntity {
		"pratyua",
		"xxxxxxx",
		baseEntity,
	}

	fmt.Println(userEntity)
	fmt.Println(userEntity.username)
	fmt.Println(userEntity.password)
	fmt.Println(userEntity.getId())
	
}