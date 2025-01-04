package main

import (
	"fmt"
	"time"
)

func main() {

	//package time se utiliza para trabajar con fechas y horas

	p := fmt.Println //asignamos esta funcion a una variable

	now := time.Now() //nos da la fecha y hora actual
	p(now)

	then := time.Date(2025, 11, 12, 20, 34, 58, 651387237, time.UTC) //creamos una fecha y hora
	p(then)

	p(then.Year()) //nos da el año

	then = then.Add(time.Hour) //agregamos una hora
	p(then)

	p(then.Before(now)) //comparamos si la fecha then es antes de la fecha now

	p(then.After(now)) //comparamos si la fecha then es despues de la fecha now

	p(then.Equal(now)) //comparamos si la fecha then es igual a la fecha now

	diff := now.Sub(then) //diferencia entre las fechas
	p(diff)

	p(diff.Hours()) //diferencia en horas

	p(diff.Minutes()) //diferencia en minutos

	p(diff.Seconds()) //diferencia en segundos

	p(then.Format("2006-01-02")) //formateamos la fecha

	p(then.Format("2006-01-02 15:04:05")) //formateamos la fecha y hora

	p(then.Format("2006-01-02 15:04:05.000000")) //formateamos la fecha y hora con milisegundos

	form := "3 04 PM"                    //formato de la fecha
	t2, _ := time.Parse(form, "8 41 PM") //parseamos la fecha
	p(t2)

}
