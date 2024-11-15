package main

import (
	"encoding/json" //permite transformar una struct en json
	"fmt"
)

type User struct {
	ID       int    `json:"id,omitempty"` // con omitempty si un campo viene vacio lo omite
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	Age      uint8  `json:"age,omitempty"`
	LastName string `json:"last_name,omitempty"` // para nombrar los campos se usa etiqueta json para su representacion en formato json
}

func (u User) Display() {
	v, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println("json: ", string(v))
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {

	user := User{
		ID:       123,
		Name:     "Caupolican",
		Address:  "San Martin 7",
		Age:      38,
		LastName: "Lautaro",
	}

	fmt.Println(user)

	user.Display()
	user.SetName("Mario Roberto")
	user.Display()
	fmt.Println()

}
