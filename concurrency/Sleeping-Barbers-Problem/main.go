// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
// open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
// waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
// sleep until a new client arrives. So the rules are as follows:
//
//		- if there are no customers, the barber falls asleep in the chair
//		- a customer must wake the barber if he is asleep
//		- if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
//		  sits in an empty chair if it's available
//		- when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
//		  and falls asleep if there are none
// 		- shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is
//	      empty
//		- after the shop is closed and there are no clients left in the waiting area, the barber
//		  goes home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Dijkstra.
//
// The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
// semaphores (mutexes) is not needed.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 200
var cutDuration = 1000 * time.Millisecond
var timeOpen = 5 * time.Second

func main() {
	//seed our random number generator
	rand.Seed(time.Now().UnixNano())
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity) // buffered channel
	doneChan := make(chan bool)                      //unbuffered channel

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Blue("The shop is open for the day!")

	// add barbers
	shop.addBarber("Frank")

	// start the barbershop as goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		color.Cyan("closing shop for the day.")
		close(shop.ClientsChan) //if this channel is not closed program will run into a deadlock state because the barber go routine will be asleep ->no new client will come and shop go routine will wait for all the barber goroutine to close hence deadlock
		shop.Open = false
		for a := 1; a <= shop.NumberOfBarbers; a++ {
			<-shop.BarbersDoneChan
		}
		close(shop.BarbersDoneChan)

		color.Green("--------------------------------------------------------------------")
		color.Green("The barbershop is now closed for the day, and everyone has gone home.")
		closed <- true
	}()

	// add clients
	client := 1

	go func() {
		for {
			//get random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				color.Blue("NO client allowed now")
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				clientName := fmt.Sprintf("Client #%d", client)
				color.Green("*** %s arrives!", clientName)
				select {
				case shop.ClientsChan <- clientName:
					color.Yellow("%s takes a seat in the waiting room.", clientName)
				default:
					color.Red("The waiting room is full, so %s leaves.", clientName)
				}
				client++
			}
		}
	}()

	// blocked until the barbershop is closed
	<-closed
}
