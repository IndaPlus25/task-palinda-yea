package main

import (
	"fmt"
	"time"
)

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", 60*time.Second)
}

func Remind(text string, delay time.Duration) {
	for {
		time.Sleep(delay)
		hour, minute, second := time.Now().Clock()
		fmt.Printf("The time is %v.%v.%v: %v\n", hour, minute, second, text)
	}
}
