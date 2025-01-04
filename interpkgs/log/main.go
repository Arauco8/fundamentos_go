package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	log.Println("Este es un mensaje de log")
	//log.Fatal("Este es un mensaje de log fatal")
	//log.Panic("Este es un mensaje de log panic")

	date := time.Now()
	file, err := os.Create(fmt.Sprintf("log_%d.log", date.UnixNano()))
	if err != nil {
		log.Panic(err.Error())
	}

	//para imprimir por consola en lugar de file pongo os.Stdout
	l := log.New(file, "Success: ", log.LstdFlags|log.Lshortfile) // log.New crea un nuevo logger
	l.Println("Este es un mensaje de log personalizado")
	l.Println("Este es otro mensaje de log personalizado")

	l.SetPrefix("Error: ")
	l.Println("Este es un mensaje de log personalizado")
	l.Println("Este es otro mensaje de log personalizado")

}
