package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	cache := NewLRUCache(3)
	//val := cache.Get(1)
	//fmt.Println(val)

	cache.Set(1, "val1")
	cache.Set(2, "val2")

	fmt.Println(cache.top.key)
	fmt.Println(cache.end.key)
	cache.Get(1)
	fmt.Println(cache.top.key)
	fmt.Println(cache.end.key)
	cache.Set(3, "val3")
	fmt.Println(cache.top.key)
	fmt.Println(cache.end.key)

	cache.Set(4, "val4")
	fmt.Println(cache.top.key)
	fmt.Println(cache.end.key)

	fmt.Println(cache.Get(2))
}
