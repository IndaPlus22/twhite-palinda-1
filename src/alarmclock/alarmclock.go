package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		t := time.Now().Format("15:04:05")
		fmt.Printf("The time is %s: %s\n", t, text)
		time.Sleep(delay)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	go Remind("Time to sleep", 1*time.Minute)
	select {}
}
