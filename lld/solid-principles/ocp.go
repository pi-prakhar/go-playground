package main

import "fmt"

/*
Open/Closed Principle (OCP)
Software entities (e.g., classes, modules) should be open for extension but closed for modification.

In Go:

1) Use interfaces to define behavior and allow extensibility.
2) Write code that can be extended with new functionality without modifying existing code.
*/

type PaymentProcessor interface {
	Pay(amount float64) error
}

type CreditCardProcessor struct{}

func (ccp *CreditCardProcessor) Pay(amount float64) error {
	fmt.Println("Paid with credit card:", amount)
	return nil
}

type PayPalProcessor struct{}

func (pp *PayPalProcessor) Pay(amount float64) error {
	fmt.Println("Paid with PayPal:", amount)
	return nil
}

func ProcessPayment(p PaymentProcessor, amount float64) {
	p.Pay(amount)
}
