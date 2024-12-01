package main

import (
	"fmt"
	"vending-machine/payment"
)

type Item struct {
	ID    int
	name  string
	price int
	count int
}

type VendingMachine struct {
	slotMap      map[int]*Item
	changeTotal  int
	changeCoins  map[int]int
	isFunctional bool
}

func NewVendingMachine(slotMap map[int]*Item, changeTotal int, changeCoins map[int]int, isFunctional bool) *VendingMachine {
	return &VendingMachine{
		slotMap:      slotMap,
		changeTotal:  changeTotal,
		changeCoins:  changeCoins,
		isFunctional: isFunctional,
	}
}

func (v *VendingMachine) GetPrice(itemID int) (int, error) {
	if item, ok := v.slotMap[itemID]; ok {
		return item.price, nil
	}
	return 0, fmt.Errorf("No such Item")
}

func (v *VendingMachine) GetAvailability(itemID int) bool {
	if item, ok := v.slotMap[itemID]; ok && item.count > 0 {
		return true
	}
	return false
}

func (v *VendingMachine) BuyItem(itemID int, amount int, paymentMethod payment.PaymentMethod) (string, bool, int) {
	// if available
	if item, ok := v.slotMap[itemID]; ok && item.count > 0 {
		// if amount < item price -> message, false, refund the amount
		if amount < item.price {
			return "Item price is more than the paid amount", false, amount
			// else if change > 0 && change > change total -> message, false, refund the amount
		} else if (amount-item.price) > 0 && (amount-item.price) > v.changeTotal {
			return "No change left in vending machine, Try Later !!!", false, amount
			// else -> do payment
		} else {
			// if success -> reduce count -> reduce change -> message, true, refund the amount
			if paymentMethod.Pay() {
				change := amount - item.price
				v.changeTotal -= change
				item.count -= 1
				return "Success", true, change
				// else -> message , false, refund the amount
			} else {
				return "Payment Failed", false, amount
			}
		}
	} else {
		// else return message , false, refund the amount
		return "Wrong ID / Item not available", false, amount
	}
}

func main() {
	item1 := Item{
		ID:    1,
		name:  "name1",
		price: 10,
		count: 2,
	}

	item2 := Item{
		ID:    2,
		name:  "name2",
		price: 20,
		count: 2,
	}
	slots := make(map[int]*Item)
	coins := make(map[int]int)
	coins[1] = 10
	slots[item1.ID] = &item1
	slots[item2.ID] = &item2
	vm := NewVendingMachine(slots, 10, coins, true)
	vm.GetAvailability(1)
	vm.GetPrice(1)
	vm.GetAvailability(3)
	vm.GetPrice(3)

	cash := payment.Cash{
		Denomination: map[int]int{
			10: 2,
		},
		Total:    20,
		Currency: "INR",
	}
	msg, success, refund := vm.BuyItem(2, cash.Total, &cash)
	fmt.Println(msg, success, refund)
	msg, success, refund = vm.BuyItem(2, cash.Total, &cash)
	fmt.Println(msg, success, refund)
	msg, success, refund = vm.BuyItem(2, cash.Total, &cash)
	fmt.Println(msg, success, refund)
}
