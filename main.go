package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Hello Tracee!")
	for i := 0; i < 10; i++ { // Need to wait until Tracee is ready
		fmt.Printf("Tick %d\n", i)
		time.Sleep(time.Second)
	}
	doSomethingMalicious()
}

func doSomethingMalicious() {
	cmd := exec.Command("./poc.py")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("exploit failed, ", err)
	}
	fmt.Println("exploit successful!")
}
