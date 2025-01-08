package main

import (
	"fmt"
	"time"
)

func main() {
	go myProcess("A")
	go myProcess("B")
	time.Sleep(3 * time.Second)

	//si el programa termina, las goroutines también terminan no importa si no terminaron

	//channels

	myFirstChannel := make(chan string) //crea un canal de tipo string con la palabra reservada chan
	go myProcessWithChannel("C", myFirstChannel)

	result := <-myFirstChannel
	fmt.Println(result)
	close(myFirstChannel) //cierra el canal
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok" //envia el valor de p al canal, asigna un valor al canal
}

func myProcess(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
