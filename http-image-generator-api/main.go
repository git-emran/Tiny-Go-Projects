package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/core"
)

type jobData struct {
	Name    string
	Message string
}

func (j *jobData) Bytes() []byte {
	fmt.Printf("%s:%s\n", j.Name, j.Message)
	res := sleepSomeTime()
	j = &jobData{Name: "Iam Awake", Message: res}
	b, _ := json.Marshal(j)
	return b
}

func sleepSomeTime() string {
	seconds := rand.Intn(20)
	sleepTime := time.Duration(seconds) * time.Second
	time.Sleep(sleepTime)
	return fmt.Sprintf("Commander, I slept: %d seconds", seconds)
}

func main() {
	rand.Intn(100)
	taskN := 100
	rets := make(chan string, taskN)
	q := queue.NewPool(30, queue.WithFn(func(ctx context.Context, m core.TaskMessage) error {
		var data jobData
		json.Unmarshal(m.Bytes(), &data)
		rets <- "Hello, " + data.Name + ", " + data.Message
		return nil
	}))
	defer q.Release()

	for i := range taskN {
		go func(i int) {
			q.Queue(&jobData{
				Name:    "Sleeping Gophers",
				Message: fmt.Sprintf("hello commander, I am handling the job: %d", +i),
			})
		}(i)
	}

	for range taskN {
		fmt.Println("message: ", <-rets)
		time.Sleep(10 * time.Millisecond)
	}
}
