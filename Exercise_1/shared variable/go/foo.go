package main

import (
	"fmt"
	"runtime"
	"sync"
)

type command struct {
	action string
	result chan int
}

func numberServer(commands chan command, done chan bool) {
	var i int
	for {
		select {
		case c := <-commands:
			switch c.action {
			case "increment":
				i++
			case "decrement":
				i--
			case "get":
				c.result <- i
			}
		case <-done:
			return
		}
	}
}

func worker(action string, commands chan command, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 1000000; j++ {
		commands <- command{action: action}
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	// Opprett kanaler
	commands := make(chan command)
	done := make(chan bool)

	// Start numberServer
	go numberServer(commands, done)

	var wg sync.WaitGroup
	wg.Add(2)

	// Start gorutiner for inkrementering og dekrementering
	go worker("increment", commands, &wg)
	go worker("decrement", commands, &wg)

	// Vent på at begge gorutiner fullfører
	wg.Wait()

	// Få den endelige verdien av i
	result := make(chan int)
	commands <- command{action: "get", result: result}
	finalValue := <-result

	fmt.Println("The magic number is:", finalValue)

	// Avslutt numberServer
	done <- true
}

