package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(msg string, c chan string) {
	fmt.Println("hello, ",msg)

	for i := 0; ; i++ {
		c <- fmt.Sprintf("hello %s ! %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	//go say("Captain")
	fmt.Println("Concurrency")

	c := make(chan string)
	//go boring("boring!", c)
	go say("Captain",c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You day: %q\n", <-c) }
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}