package app

type BillUser struct {
	ID int
	Bill
	User
	Price float64
	Paid  bool
}
