package models

type Employee struct {
	Id   string `json:"id,omitempty"           bson:"_id,omitempty"`
	Name string `json:"name,omitempty"         bson:"name,omitempty"`
	Age  int    `json:"age,omitempty"          bson:"age,omitempty"`
}


