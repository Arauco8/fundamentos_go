package main

import (
	"errors"
	"fmt"

	"proyectos/fundamentos_go/errors/mypackage" //importamos el paquete mypackage
)

type MyCustomError struct {
	ID string
}

func (e MyCustomError) Error() string {
	return fmt.Sprintf("Error with ID: %s", e.ID)
}

func main() {

	var err error
	err = errors.New("error de prueba")
	fmt.Println(err)
	fmt.Println(err.Error()) //lo transforma en un string

	err2 := fmt.Errorf("error de prueba err, string: %s, number:  %d", "my string", 231)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	err3 := TestError(1)
	fmt.Println(errors.As(err3, &MyCustomError{})) //nos dice si el error es de tipo MyCustomError
	fmt.Println()
	err4 := TestError(4)
	fmt.Println(errors.As(err4, &MyCustomError{})) //nos dice si el error es de tipo MyCustomError
	fmt.Println()
	err5 := errors.Join(err, err2, err3) //nos permite concatenar errores
	fmt.Println(err5)
	fmt.Println()
	fmt.Println(errors.Is(err5, err2)) //nos dice si el error err2 esta contenido en err5
	fmt.Println(errors.Is(err5, err4))
	fmt.Println()
	err6 := fmt.Errorf("error1: [%w]", err) //nos permite anidar un error
	err7 := fmt.Errorf("error2: [%w]", err6)
	fmt.Println(err7)
	fmt.Println()
	fmt.Println(errors.Unwrap(err7)) //nos permite desanidar un error
	fmt.Println(errors.Unwrap(errors.Unwrap(err7)))
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err7))))
	fmt.Println()

	defer func() { //ejecuta al final del programa. Utilizamos el defer para recuperar el panic (errores)
		fmt.Println("end main")
		r := recover() //nos permite recuperar el panic, nos permite manejar errores que se van generando
		if r != nil {
			fmt.Println("recovered from panic", r)
		}
	}()

	/*v := 0
	_ = 5 / v*/

	mypackage.Run()
	//panic("something bad happened") //genera un panic, forzamos un error. Es importante que el defer este antes que el panic o cualquier error que pueda llegar a pasar

	//fmt.Println("end of main")

}

func TestError(v int) error {
	if v == 1 {
		return MyCustomError{ID: "123"}
	} else {
		return errors.New("can't work with 42")
	}
}
