package main

import (
	"fmt"
	"go-goroutine/utils"
	"strconv"
	"testing"
	"time"
)

func TestCreateGoRoutine(t *testing.T) {
	go utils.RunHelloWorld() //go routine sangat ringan
	fmt.Println("Ups")

	time.Sleep(2 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go utils.DisplayNumber(i)
	}
	time.Sleep(2 * time.Second)
}

func TestCobaChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Arfian" //data dikirim ke channel
	}()

	data := <-channel //mengambil data dari channel
	fmt.Println(data)
	close(channel) //Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak
}

func TestChannelSebagaiParameter(t *testing.T) {
	channel := make(chan string)
	go utils.GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	close(channel)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	go utils.OnlyIn(channel)
	go utils.OnlyOut(channel)
	time.Sleep(2 * time.Second)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan int, 3)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	go func() {
		for {
			i := <-channel
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		channel <- i
	}
	close(channel)

}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel { //otomatis berhenti ketika channel di close()
		fmt.Println(data)
	}

}
