package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepSomeTime() string {
	sleepTime := time.Duration(rand.Intn(60)) * time.Second
	message := fmt.Sprintf("%s\n", sleepTime)
	time.Sleep(sleepTime)
	return message
}

func job(i int, rets chan string) {
	sleepSomeTime()
	rets <- fmt.Sprintf("Hi gopher, handle the job: %02d", +i)
}

func main() {
	taskN := 100
	rets := make(chan string, taskN)
	for i := 0; i < taskN; i++ {
		go job(i, rets)
	}

	for i := 0; i < taskN; i++ {
		fmt.Println("message:", <-rets)
		time.Sleep(20 * time.Millisecond)
	}
}
