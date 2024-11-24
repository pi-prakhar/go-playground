package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "P1", leftFork: 4, rightFork: 0},
	{name: "P2", leftFork: 0, rightFork: 1},
	{name: "P3", leftFork: 1, rightFork: 2},
	{name: "P4", leftFork: 2, rightFork: 3},
	{name: "P5", leftFork: 3, rightFork: 4},
}

type Order struct {
	mu   *sync.Mutex
	list *[]string
}

var list = make([]string, len(philosophers))
var order = Order{
	mu:   &sync.Mutex{},
	list: &list,
}

var forks = make(map[int]*sync.Mutex)
var hunger = 1
var eatTime = 0 * time.Second

func main() {
	//create order list
	// list := make([]string, len(philosophers))
	// order := Order{
	// 	mu:   &sync.Mutex{},
	// 	list: &list,
	// }
	//create forks
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	fmt.Println("Dinning Start")
	fmt.Println("______________________________________________")

	//create a wait group
	wg := &sync.WaitGroup{}
	//create a settled wait group
	settled := &sync.WaitGroup{}

	//adding count to philosophers to wait groups
	wg.Add(len(philosophers))
	settled.Add(len(philosophers))

	//start dinning
	for _, philosopher := range philosophers {
		go philosopherEats(philosopher, forks, wg, settled, &order)
	}

	//wait for all philosophers to finish eating (all goroutined to finish)
	wg.Wait()

	fmt.Println("Dinning Ends")
	fmt.Println("______________________________________________")

	fmt.Println("Order of Dinners")
	for _, philosopher := range *order.list {
		fmt.Print(philosopher)
	}

}

func philosopherEats(philosopher Philosopher, forks map[int]*sync.Mutex, wg *sync.WaitGroup, settled *sync.WaitGroup, order *Order) {
	defer wg.Done()
	fmt.Printf("%s has settled at Dinning Table\n", philosopher.name)
	//philosopher is seated
	settled.Done()
	//waiting for others to be seated as well
	settled.Wait()
	//while hungry
	for hunger > 0 {

		if philosopher.leftFork > philosopher.rightFork {
			//pick left fork
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has picked left fork{%d} at Dinning Table \n", philosopher.name, philosopher.leftFork)
			//pick right fork
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has picked right fork{%d} at Dinning Table \n", philosopher.name, philosopher.rightFork)
		} else {
			//pick right fork
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has picked right fork{%d} at Dinning Table \n", philosopher.name, philosopher.rightFork)
			//pick left fork
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has picked left fork{%d} at Dinning Table \n", philosopher.name, philosopher.leftFork)
		}

		//eat
		fmt.Printf("%s is eating now\n", philosopher.name)
		time.Sleep(eatTime)

		//leave left fork
		forks[philosopher.leftFork].Unlock()
		fmt.Printf("%s has put down left fork{%d} at Dinning Table\n", philosopher.name, philosopher.leftFork)

		//leave right fork
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("%s has put down right fork{%d} at Dinning Table\n", philosopher.name, philosopher.rightFork)
		hunger--
	}
	//finished eating
	order.mu.Lock()
	*order.list = append(*order.list, philosopher.name)
	order.mu.Unlock()
	//leave table
	fmt.Printf("%s has left the dinning table\n", philosopher.name)

}
