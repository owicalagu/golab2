package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	for {
		fmt.Println("foo sending ping")
		channel <- "ping"

		pong := <-channel
		fmt.Println("foo received ", pong)
	}
}

func bar(channel chan string) {
	for {
		ping := <-channel
		fmt.Println("bar recieved ", ping)

		fmt.Println("bar sending pong")
		channel <- "pong"
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	channel := make(chan string)
	go foo(channel)
	go bar(channel)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
