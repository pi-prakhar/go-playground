package main

import "fmt"

/*
Interface Segregation Principle (ISP)
Clients should not be forced to depend on interfaces they do not use.

In Go:

1) Keep interfaces small and focused.
2) Split larger interfaces into smaller, more specific ones.
*/

// Define two small interfaces
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Combine Reader and Writer into a single interface
type ReadWriter interface {
	Reader
	Writer
}

// A type that implements ReadWriter
type File struct{}

func (f File) Read(p []byte) (n int, err error) {
	fmt.Println("Reading data...")
	return len(p), nil
}

func (f File) Write(p []byte) (n int, err error) {
	fmt.Println("Writing data...")
	return len(p), nil
}

func main() {
	var rw ReadWriter = File{}
	rw.Read([]byte("hello"))
	rw.Write([]byte("world"))
}
