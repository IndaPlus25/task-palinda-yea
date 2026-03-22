package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
// Because you give the channel a value it pauses the current goroutine until the value is used,
// but since it's on the line after it never happens, causing a deadlock
// Giving the channel a value in a seperate goroutine means that the program can continue to the next line and no deadlock occurs
// Could also just give the channel a buffer
func main() {
	ch := make(chan string)
	go func() { ch <- "Hello world!" }()
	fmt.Println(<-ch)
}
