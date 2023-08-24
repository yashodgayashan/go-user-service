package main

type User struct {
	ID       int    "json:id,omitempty"
	Name     string "json:name,omitempty"
	Age      int    "json:age,omitempty"
	Location string `json:"location,omitempty"`
}

var users []User
var nextID int = 1
