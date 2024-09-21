package main

import "fmt"

/*
*
Reference --- practical-go-lessons.com/chap-22-maps
## Maps could also be implemented via structs then creating a slice out of those structs but :

1)We have to iterate potentially over all elements of the slice. Imagine that your slice contains thousands of elements! The impact on performance can be important.

2)The code written is not short. We use a for loop range and a nested comparison. Those five lines are not easy to read.
_______________________________________________________________________________________________________________________________________
## What is a map?
A map is an unordered collection of elements of type T that are indexed by unique keys of type U.
An element of a map is called a “map entry”. It’s also usually named a key-value pair

_______________________________________________________________________________________________________________________________________

## Backend Logic for maps (General)

1) Maps are implemented using hashTables.
2) Hash Table has 2 components.
	-> Hash Function - its role is to transform a key into a unique identifier. For instance, the key 1930 will be passed to the hash function, and it will return “4”.
	-> An indexed storage - that is used to keep the values in memory. The storage is eventually organized in buckets. Each bucket can store a specific number of values.
3) The following time complexity applies in general for hash tables :
	-> Insertion : O(1)
	-> Search : O(1)

## Backend Logic for maps (GO Lang)
1) The source code is located in the runtime package (runtime/map.go).
	A Go map is an array of “buckets”
	A bucket contains a maximum number of 8 key/element pairs (also called 8 entries).
	Each bucket is identified by a number (an id).

2) rest is same logic -> key -> hashFunction(key) -> returns hashValue (h) -> h -> transformed -> bucket id -> id used to find value for given key

3) In Golang you cannot access the memory address of any value by giving a key.
REFER - for futher lookup
**/

func main() {
	fmt.Println("Welcome to session on maps")

	//create map using make
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	fmt.Println(languages)

	//create map using 'map literal'
	worldCupWinners := map[int]string{
		1930: "Uruguay",
		1934: "Italy",
		1938: "Italy",
		1950: "Uruguay"}
	fmt.Println(worldCupWinners)

	//get value from map
	fmt.Println("Js is ", languages["JS"])

	//delete value from map
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//iterate over a map
	for key, value := range languages {
		fmt.Println(key, " : ", value)
	}

	//fetch value of map by checking
	if val, ok := languages["JS"]; ok {
		fmt.Println(val)
	}

}
