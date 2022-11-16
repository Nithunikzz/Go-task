package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go wait(ch)

	val := <-ch
	fmt.Println(val)

	intCh := make(chan int, 10)
	// go getValues(intCh)
	// for val := range intCh {
	// 	fmt.Printf("val: %d\n", val)
	// }

	t := time.NewTicker(500 * time.Millisecond)
	go getValuesPeriodically(intCh, t)
	go func() {
		time.Sleep(20 * time.Second)
		t.Stop()
	}()

	for {
		select {
		case valInt := <-intCh:
			fmt.Printf("Recevied: %d\n", valInt)
		case <-time.After(10 * time.Second):
			fmt.Println("waited for 10 seconds after last message, returning")
			return
		}
	}
}

func wait(ch chan bool) {
	time.Sleep(5 * time.Second)
	ch <- true
}

func getValues(ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		fmt.Printf("wrote %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func getValuesPeriodically(ch chan int, t *time.Ticker) {
	i := 0
	for {
		select {
		case <-t.C:
			ch <- i
			fmt.Printf("wrote %d\n", i)
			i++
		case <-time.After(5 * time.Second):
			fmt.Println("No ticker for 5 seconds, returning from getValuesPeriodically")
			return
		}
	}
}
