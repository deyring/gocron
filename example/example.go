package main

import (
	"fmt"

	"github.com/deyring/gocron"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	fmt.Println("creating new scheduler")
	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	for _, j := range &s.Jobs {
		if j != nil {
			fmt.Printf("%v \n", j.JobFunc)
		}
	}
	fmt.Println("Scheduler started")
	<-s.Start()
}
