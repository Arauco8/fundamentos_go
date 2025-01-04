package main

import (
	"fmt"
	"os"
)

func main() {
	//package os utilizar funciones del sistema operativo

	file, error := os.Open("main.go")
	if error != nil {
		fmt.Println(error)
		os.Exit(1) //termina la ejecución del programa
	}

	fmt.Println(file)
	v, _ := file.Stat()
	fmt.Println(v.Name(), v.Size(), v.ModTime())

	data := make([]byte, 1024) //crear un slice de bytes

	c, error := file.Read(data) //leer el archivo
	if error != nil {
		fmt.Println(error)
		os.Exit(1) //termina la ejecución del programa
	}
	fmt.Printf("read: %d bytes: %q \n", c, data[:c])
	fmt.Println(string(data[:c]), c)
	fmt.Println()
	fmt.Println(os.Getenv("HOME"))   //obtener la variable de entorno HOME
	fmt.Println(os.Getenv("MI_ENV")) //obtener la variable de entorno MI_ENV
	os.Setenv("MI_ENV", "valor")     //establecer una variable de entorno
	fmt.Println(os.Getenv("MI_ENV")) //obtener la variable de entorno MI_ENV

	dir, _ := os.Getwd() //obtener el directorio de trabajo actual
	fmt.Println(dir)

	os.Setenv("ENV_EXISTS", "")
	fmt.Println(os.Getenv("ENV_EXISTS"))
	fmt.Println(os.Getenv("ENV_NOT_EXISTS"))

	env, ok := os.LookupEnv("ENV_EXISTS") //buscar una variable de entorno
	fmt.Println(env, ok)

	env1, ok := os.LookupEnv("ENV_NOT_EXISTS")
	fmt.Println(env1, ok)

	os.Setenv("DB_USERNAME", "root0")
	os.Setenv("DB_PASSWORD", "root123")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "test")

	dbURL := os.ExpandEnv("${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}")
	fmt.Println(dbURL)
}
