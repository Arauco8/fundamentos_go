package main

import (
	"errors"
	"fmt"
)

func main() {

	var err error
	err = errors.New("error de prueba")
	fmt.Println(err)
	fmt.Println(err.Error()) //lo transforma en un string

	err2 := fmt.Errorf("error de prueba err, string: %s, number:  %d", "my string", 231)
	fmt.Println(err2)
	fmt.Println(err2.Error())

}
