package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ 
	sync.Mutex 
}

type Philosopher struct {
	id              int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philosopher) eatProcess() {
	if p.id%2 == 0 {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
	} else {
		p.rightChopstick.Lock()
		p.leftChopstick.Lock()
	}

	fmt.Println("Philosopher",p.id,"grab the chopsticks");

	fmt.Printf("Philosopher %d is eating\n", p.id)
	time.Sleep(5 * time.Second) // Simulate eating time

	p.leftChopstick.Unlock()
	p.rightChopstick.Unlock()

	fmt.Println("Philosopher",p.id,"release the chopsticks");
	
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(5 * time.Second) // Simulate thinking time
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.eatProcess()
	}
}

func main() {
	var wg sync.WaitGroup

	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(&wg)
	}

	wg.Wait()
}
