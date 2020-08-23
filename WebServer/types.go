package main

import (
	"encoding/json"
	"net/http"
)

//Middleware type that assigns a label to a function that accepts and receives http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

//User struct that contains user data
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

//MetaData stores data
type MetaData interface{}

//ToJSON Converts user ToJson
func (user *User) ToJSON() ([]byte, error) {
	return json.Marshal(user)
}
