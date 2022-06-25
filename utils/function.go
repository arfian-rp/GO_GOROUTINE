package utils

import (
	"fmt"
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
