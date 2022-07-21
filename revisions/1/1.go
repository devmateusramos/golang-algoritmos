package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(phrase string) {
	fmt.Println(phrase)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Second)
}

func main() {
	go hello("goroutine1")
	go hello("goroutine2")
	time.Sleep(time.Second)
	fmt.Println("Chamada normal")
}
