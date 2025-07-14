package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)
	fanin := make(chan int)

	go send(even, odd, quit)
	go rec(even, odd, quit, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("------About end------")

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for a := range c2 {
		fmt.Println(a)
	}
	fmt.Println("------About end------")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 5 {
			break
		}

	}
	fmt.Println("------About end------")

	ch1 := make(chan int)
	q1 := make(chan int)

	go puts(ch1, q1)
	pull(ch1, q1)
}

func puts(c, q chan<- int) {
	var w1 sync.WaitGroup
	for j := 0; j < 10; j++ {
		w1.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				c <- i * j
			}
			w1.Done()
		}()

	}
	w1.Wait()
	q <- 1
	close(c)

}

func pull(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

// send channel
func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func rec(even, odd, quit <-chan int, fanin chan<- int) {
thisloop:
	for {
		select {

		case v := <-odd:
			wg.Add(1)
			go func(v int) {
				fanin <- v
				wg.Done()
			}(v)

		case v := <-even:
			wg.Add(1)
			go func(v int) {
				fanin <- v
				wg.Done()
			}(v)

		case <-quit:
			break thisloop
		}
	}
	wg.Wait()
	close(fanin)
}

func populate(c1 chan int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1, c2 chan int) {
	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			c2 <- timeConsumingWork(v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

func gen(ctx context.Context) <-chan int {
	n := 0
	dst := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
