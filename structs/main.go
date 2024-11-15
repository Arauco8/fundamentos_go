package main

import (
	"encoding/json" //permite transformar una struct en json
	"fmt"
	"proyectos/fundamentos_go/structs/commerce"
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

	p1 := commerce.Product{
		Name:  "Heladera Premium",
		Price: 200000,
		Type: commerce.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "freezer", "premium", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Naranja",
		Price: 50,
		Type: commerce.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "fruta"},
		Count: 20,
	}

	p3 := commerce.Product{
		Name:  "Cortina",
		Price: 6000,
		Type: commerce.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	cart := commerce.NewCart(11312)
	cart.AddProducts(p1, p2, p3)

	fmt.Println("PRODUCTS CART")
	fmt.Println("Total Products: ", len(cart.Products))
	fmt.Printf("Total Cart: %.2f\n", cart.Total())
	fmt.Println()
}
