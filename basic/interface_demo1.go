package main

import _ "fmt"

type Autentication interface {
	authenticated(username string) bool
}

type Permanent struct {
	username string
}

func (p Permanent) authenticated(username string) bool {
	return p.username == "tir"
}

type Outsource struct {
	username string
}

func (p Outsource) authenticated(username string) bool {
	return p.username == "yee"
}

func IsAuthenticated(username string, authen Autentication) bool {
	return authen.authenticated(username)
}
