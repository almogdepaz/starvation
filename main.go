package main

import (
	"fmt"
	"math"
	"time"
)

const count = 100000000

func main() {
	go foo1()
	for i := 0; i < 100; i++ {
		go foo()
	}
	time.Sleep(3)
	for i := 0; i < 100; i++ {
		go foo()
	}
	time.Sleep(10 * time.Second)
}

func wait() {
	ch := make(chan struct{})
	ti := time.NewTimer(1 * time.Second)
	fmt.Println("sleep")
	ts := time.Now()
	select {
	case <-ti.C:
		fmt.Println("wakeup ", time.Now().Sub(ts))
		break // keep going
	case <-ch:
		// closed while waiting the delta
		return
	}

}

func foo1() {
	for i := 0; i < count; i++ {
		wait()
	}
}

func foo() int {
	var a int
	for i := 0; i < count; i++ {
		a = int(math.Sqrt(float64(i)))
	}
	return a
}
