package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)
		for {
			//if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan
			//first parameter is the client from channel that get consumed one by one
			//second parameter is the bool value that returns false once the channel is empty and the channel is closed both conditions should be met

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				//cut hair
				color.Green("%s is cutting %s's hair.", barber, client)
				time.Sleep(shop.HairCutDuration)
				color.Green("%s is finished cutting %s's hair.", barber, client)
			} else {
				//shop is closed, so send the barber home and close this go routine
				color.Cyan("%s is going home.", barber)
				shop.BarbersDoneChan <- true
				return
			}
		}
	}()
}
