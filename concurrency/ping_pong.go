package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}



func main() {
	table := make(chan *Ball)

	go play("Anime",table)
	go play("Banana",table)

	ball := new(Ball)
	table <- ball
	time.Sleep(3 * time.Second)
	<-table //ส่งออกมาที่ Main
}

func play(name string,table chan *Ball) {
	for {
		ball := <- table
		ball.hits++

		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)

		table <- ball
	}
}
