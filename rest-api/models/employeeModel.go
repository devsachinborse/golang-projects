package models

type Employee struct {
	ID         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	Department string `json:"department" bson:"department"`
}
