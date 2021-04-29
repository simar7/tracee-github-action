package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Hello Tracee!")
	for i := 0; i < 1; i++ {
		fmt.Printf("Tick %d\n", i)
		time.Sleep(time.Second)
	}
	doSomethingMalicious()
}

func doSomethingMalicious() {
	exec.Command("./poc.py").Run()
}
