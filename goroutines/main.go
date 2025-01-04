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
}

func myProcess(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
