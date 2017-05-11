package main

import (
	"net/http"
)

type User struct {
	Name  string
	Phone string
}

type Victim struct {
	UserId            string
	Combination       []int
	CombinationLength int
	Token             string
	FirebaseToken     string
}

const db = "test"
const collection = "victims"
const db_addr = "127.0.0.1:27017"

func main() {
	http.HandleFunc("/", fetchUsers)
	http.ListenAndServe(":8080", nil)
}
