package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	//goroutine1()	//Will run both goroutines concurrently to completion 
	//goroutine2()	//Will run both goroutines until main finishes
	//goroutine3() 	//Will run both goroutines until user presses enter
	//goroutine4()	//Will run a goroutine until the waitGroup is done
	//goroutine5()	//Will run a goroutine until the msg is received over the channel
	//goroutine6()	//Will run a goroutine until the channel is closed
	//goroutine7()	//Will run a goroutine until the channel is closed
	//goroutine8()	//Will run a goroutine until the channel has delivered two messages
	//goroutine9()	//Will run a goroutine until the channel has delivered two messages
	goroutine10()	//Will run a goroutine until the first 100 fibonacci are delivered
}

func goroutine1() {
	go count("sheep") 
	count("fish")	
}

func goroutine2() {
	go count("sheep") 
	go count("fish")	
}

func goroutine3() {
	go count("sheep") 
	go count("fish")
	fmt.Scanln()
}

func goroutine4() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1) //essentially a counter, incremented by the number of goroutines in the group
	go func () { //using an anonymous wrapper func
		count("sheep") 
		waitGroup.Done()
	}()

	waitGroup.Wait() //will block until counter is zero
}

func goroutine5() {
	c := make(chan string) //Channels are useful to communicate across different goroutines 
	go countChannel("sheep", c)

	msg := <- c // msg will receive the output of the channel (Note: this is a blocking function)
	fmt.Println(msg)
}

func goroutine6() {
	c := make(chan string) //Channels are useful to communicate across different goroutines 
	go countChannel("sheep", c)

	for {
		msg, open := <- c  // msg will receive all output from the channel in a loop
		if !open { //the loop will break when the channel is no longer open
			break
		}
		fmt.Println(msg)
	}
}

func goroutine7() {
	c := make(chan string) //Channels are useful to communicate across different goroutines 
	go countChannel("sheep", c)

	for msg := range c { // msg will receive all output from the channel in a loop
		fmt.Println(msg)
	} //the loop will break automatically when the channel is no longer open
}

func goroutine8() {
	c := make(chan string, 2) //Note: the second param is the number of inputs to be carried across the channel before closing
	c <- "hello"
	c <- "world"
	//c <- "third message" //If a 3rd message is added to the channel of size 2, then deadlock error will occur

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}

func goroutine9() {
	fast := make(chan string) 
	slow := make(chan string) 

	go func() {
		for {
			fast <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			slow <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for i := 0; i < 5; i++{
		fmt.Println(<- fast) 
		fmt.Println(<- slow) //Because reading from the channel is a blocking call, the output will only once the slow channel has received it's input
	}

	for {
		select { //With a select statement, Println will not be blocked and outputs will be run whenever the channel has received it's input
		case msg1 := <- fast:
			fmt.Println(msg1) 
		case msg2 := <- slow:
			fmt.Println(msg2) 
		}
	}
}


func goroutine10() {
	jobs := make(chan int, 100) //buffered channel
	results := make(chan int, 100)

	go worker(jobs, results) //this call can be duplicated to utilize a multi-core processor

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results) //TODO: Fix integer overflow on fib numbers larger than int max
	}

}

//<-chan are channels that are only ever receiving
//chan<- are channels that are only ever sending 
func worker(jobs <-chan int, results chan<- int){
	fib_map := make(map[int]int) // Initialized as map[key]value

	for n := range jobs {
		results <- fib(n, fib_map)
	}0
}

func fib(n int, fib_map map[int]int) int {
	nth_fib, exists := fib_map[n]
    if exists {
        return nth_fib
    } else {
		if n <= 1 {
			fib_map[n] = n
			return n
		}
		fib_map[n] = fib(n - 1, fib_map) + fib(n - 2, fib_map)
		return fib_map[n]
	}
}

func count(thing string) {
	for i:= 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countChannel(thing string, c chan string) {
	for i:= 0; i < 5; i++ {
		c <- thing //send the string value through the channel
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}


