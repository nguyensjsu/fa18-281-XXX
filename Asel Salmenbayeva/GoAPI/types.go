package main

type payment struct {
	PaymentID   string
	Amount     float32       `bson:"Amount" json:"Amount"`

}
