package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type person struct {
	First string
	Last  string
}

func (p *person) Say() {
	fmt.Printf("My name is %s %s\n", p.First, p.Last)
}

type human interface {
	Say()
}

func saySomething(h human) {
	h.Say()
}

func main() {

	p1 := person{"Yuriy", "Pavlov"}
	p1.Say()
	saySomething(&p1)

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 10; i++ {
			runtime.Gosched()
			//fmt.Println("Goroutine 1", i)
		}
		fmt.Println("Goroutine 1")
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			runtime.Gosched()
			//fmt.Println("Goroutine 2", i)
		}
		fmt.Println("Goroutine 2")
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			runtime.Gosched()
			//fmt.Println("Goroutine 3", i)
		}
		fmt.Println("Goroutine 3")
		wg.Done()
	}()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

	var counter int64
	gc := 100
	wg.Add(gc)
	//var mu sync.Mutex

	for j := 0; j < gc; j++ {
		go func() {
			/*
				mu.Lock()
				runtime.Gosched()
				var i int64
				i = counter
				i++
				counter = i
				mu.Unlock()
			*/

			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(counter)
}
