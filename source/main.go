package main

import (
	"log"
	"os"
	"time"
)

func main() {
	input := NewInput(os.Stdin)

	T := input.Int()

	startTime := time.Now()
	for i := 1; i <= T; i++ {
		runTestCase(input, i)
	}
	totalTime := time.Now().Sub(startTime)

	log.Println("Total time:", totalTime)
}

func runTestCase(input *Input, i int) {
	warningTimer := time.NewTimer(1 * time.Second)

	doneChan := make(chan bool)

	go func() {
		testCase(input, i)

		doneChan <- true
	}()

loop:
	for {
		select {
		case <-warningTimer.C:
			log.Println("Long calculation")
		case <-doneChan:
			break loop
		}
	}
}
