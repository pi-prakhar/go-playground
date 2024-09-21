package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTimeLocal := time.Now().Local()
	presentTimeUTC := time.Now().UTC()
	fmt.Println(presentTimeLocal)
	fmt.Println(presentTimeLocal.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println(presentTimeUTC)
	fmt.Println(presentTimeUTC.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.December, 24, 24, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
