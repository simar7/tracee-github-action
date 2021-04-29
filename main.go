package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Tracee!")
	for i := 0; i < 60; i++ {
		fmt.Printf("Tick %d\n", i)
		time.Sleep(time.Second)
	}
}
