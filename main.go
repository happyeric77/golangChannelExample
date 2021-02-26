package main

import (
	"log"
	"sync"
	"fmt"
	"time"
)
// unbuffered channel:
/* Unbuffered channel needs a reciever right after a message is emitted to the channel.
Reason: the channel is not able to buffer the message */

func unbufferedChanDemo() {
	// log.Println("=========unbuffered channel demo===========")
	var wait sync.WaitGroup

	// Create a unbuffered channel
	f := make(chan string)

	wait.Add(1)
	go func(){
		defer wait.Done()
		for i := 0; i < 5; i ++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()

	// Add a goroutine to emmit a message to channel
	wait.Add(1)
	go func(){
		defer wait.Done()
		time.Sleep(4*time.Second)
		fmt.Println("message is being emmited to channel")
		f <- "Hello message"
		time.Sleep(5*time.Millisecond)
		fmt.Println("message emmited to channel")		
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
	}()
	
	// the wait group needs to be a go routine becuase it is not possible to wait any time after the message is emmited to the channel
	go func(){
		wait.Wait()
	}()

	// the reciver here to recieve the message 
	fmt.Println("Reciever retrieved message from channel: ", <-f)
}


// buffered channel:
/* buffered channel does not need a reciever right after a message is emitted to the channel.
 Reason: the channel is able to buffer the value */

func bufferedChanDemo() {
	var wait sync.WaitGroup

	// Create a buffered channel
	f := make(chan string, 50)

	wait.Add(1)
	go func(){
		defer wait.Done()
		for i := 0; i < 5; i ++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()

	// Add a goroutine to emmit a message to channel
	wait.Add(1)
	go func(){
		defer wait.Done()
		time.Sleep(4*time.Second)
		fmt.Println("message is being emmited to channel")
		f <- "Hello message"
		time.Sleep(5*time.Millisecond)
		fmt.Println("message emmited to channel")
	}()

	wait.Add(1)
	go func() {
		defer wait.Done()
	}()
	
	// the wait group does not need to be a go routine becuase the emmited messaged can be buffered into the channel
	wait.Wait()

	// the reciver here to recieve the message 
	fmt.Println("Reciever retrieved message from channel: ", <-f)
}

func main() {
	log.Println("=========unbuffered channel demo===========")
	unbufferedChanDemo()

	// log.Println("===========buffered channel demo===========")
	// bufferedChanDemo()
}
