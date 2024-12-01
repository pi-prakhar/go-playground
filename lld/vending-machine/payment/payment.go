package payment

type PaymentMethod interface {
	Pay() bool
}

type UPI struct {
	ID    string
	Total int
}

func (upi *UPI) Pay() bool {
	return true
}

type Cash struct {
	Denomination map[int]int
	Currency     string
	Total        int
}

func (c *Cash) Pay() bool {
	return true
}
