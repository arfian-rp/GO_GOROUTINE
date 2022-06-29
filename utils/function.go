package utils

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello WOrlk")
}

func DisplayNumber(number int) {
	fmt.Println("display", number)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Ini Response"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "data string"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func ShowInfoCPU() {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
}
