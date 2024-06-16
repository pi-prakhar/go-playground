package main

import (
	"time"
)

/**
Process:
A process is an instance of a program that is currently executed by a computer.
ps xau


Thread:
A thread represents the sequential execution of a set of instructions. A single process can create multiple threads. For each process we have at least one thread of execution. By creating more than one thread we create multiple execution streams that may share some data.

Concurrency:
“Concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the final outcome”1 Concurrency refers to an ability. Go supports concurrency. We can write a Go program that will run a set of tasks concurrently.

Parallelism:
The previous notion of concurrency can be mixed up with parallelism. “Concurrency is not parallelism” (Rob Pike) 2. Parallelism refer to tasks that are executed simultaneously (at the same time). A concurrent program might run tasks in parallel.

**/

func main() {
	Base()
	WaitGroup()
	time.Sleep(1 * time.Second)
}
