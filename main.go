package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Hello Tracee!")
	for i := 0; i < 5; i++ {
		fmt.Printf("Tick %d\n", i)
		time.Sleep(time.Second)
	}
	doSomethingMalicious()
}

func doSomethingMalicious() {
	cmd := exec.Command("./poc.py")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdout)
}
