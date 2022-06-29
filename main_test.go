package main

import (
	"fmt"
	"go-goroutine/utils"
	"runtime"
	"strconv"
	"sync"
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

func TestSelectChannel(t *testing.T) {
	//Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random

	channel1 := make(chan string)
	channel2 := make(chan string)

	counter := 0

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "data1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "data2"
	}()

	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	//Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random

	channel1 := make(chan string)
	channel2 := make(chan string)

	counter := 0

	go func() {
		time.Sleep(2 * time.Second)
		// channel1 <- "data1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		// channel2 <- "data2"
	}()

	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2:", data)
			counter++
		default:
			fmt.Println("menunggu data")
			counter += 2
		}

		if counter == 2 {
			break
		}
	}

}

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 1000; j++ {
				x += 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(x) //harusnya 1.000.000
}

func TestMutex(t *testing.T) {
	/*
		untuk mengatasi masalah tersebut di golang ada anmanya mutex.
		mutex  =>  melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock.
	*/
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 1000; j++ {

				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(x)
}

func TestRWMutex(t *testing.T) {
	//RWMutex => Read Write Mutex
	account := utils.BankAccount{}

	for i := 1; i <= 50; i++ {
		go func() {
			for j := 1; j <= 50; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(account.GetBalance())
}

func TestSimulasiDeadLock(t *testing.T) {
	u1 := utils.UserBalance{
		Name: "Yudi",
	}
	u2 := utils.UserBalance{
		Name: "Yuda",
	}

	go utils.Transfer(&u1, &u2, 1000)
	go utils.Transfer(&u2, &u1, 1000)

	time.Sleep(5 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	//Solusi DeadLock

	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go utils.RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Complete")
}

func TestOnce(t *testing.T) {
	//memastikan funcion hanya dieksekusi 1x
	counter := 0
	OnlyOnce := func() {
		counter++
	}

	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}

func TestGOMAXPROCS(t *testing.T) {
	utils.ShowInfoCPU()
	runtime.GOMAXPROCS(20) //mengubah thread
	utils.ShowInfoCPU()
}
