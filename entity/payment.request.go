package entity

type PaymentRequest struct {
	OrderID string
	Amount  int64
	FName   string
	LName   string
	Email   string
}

func NewPaymentRequest(orderID string, amount int64, fName, lName, email string) *PaymentRequest {
	return &PaymentRequest{
		OrderID: orderID,
		Amount:  amount,
		FName:   fName,
		LName:   lName,
		Email:   email,
	}
}
