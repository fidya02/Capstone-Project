package entity

type Payments struct {
	ID      int64
	OrderID string
	UserID  int64
	Amount  int64
	Status  string
}

func NewPayments(orderID string, userID int64, amount int64, status string) *Payments {
	return &Payments{
		OrderID: orderID,
		UserID:  userID,
		Amount:  amount,
		Status:  status,
	}
}
