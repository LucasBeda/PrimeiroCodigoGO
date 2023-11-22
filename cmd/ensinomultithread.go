package cmd

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	//go processando() //T2
	//go processando() //T3
	//processando()    //T1
	canal := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			canal <- i //T2
			fmt.Println("Jogou no canal ", i)
		}
	}()

	//fmt.Println(<-canal) //esvaziar o canal

	//load balancer
	go worker(canal, 1)
	worker(canal, 2)
}

func worker(canal chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal ", <-canal, "no worker", workerID)
		time.Sleep(time.Second)
	}
}
