package models

type Person struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
