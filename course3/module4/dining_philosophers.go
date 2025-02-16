package main

import (
	"fmt"
	"sync"
	"time"
)

var eatSemaphore = make(chan struct{}, 2)

type Chopstick struct {
	number int
	sync.Mutex
}

type Philosopher struct {
	number          int
	leftCS, rightCS *Chopstick
}

func (p Philosopher) eatOnce() {
	cs1, cs2 := p.leftCS, p.rightCS
	if cs1.number > cs2.number {
		cs1, cs2 = cs2, cs1
	}
	cs1.Lock()
	cs2.Lock()
	eatSemaphore <- struct{}{}
	fmt.Println("starting to eat ", p.number)
	time.Sleep(1 * time.Second)
	fmt.Println("finishing eating ", p.number)
	<-eatSemaphore
	cs1.Unlock()
	cs2.Unlock()
}

func (p Philosopher) eatTimes(times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < times; i++ {
		p.eatOnce()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	const numPhilosophers = 5
	philosophers := make([]Philosopher, numPhilosophers)
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{number: i}
	}
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = Philosopher{
			number:  i,
			leftCS:  chopsticks[i],
			rightCS: chopsticks[(i+1)%numPhilosophers],
		}
	}
	var wg sync.WaitGroup
	wg.Add(numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		go philosophers[i].eatTimes(3, &wg)
	}
	wg.Wait()
}
