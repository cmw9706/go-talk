package main

//Person Model
type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}
