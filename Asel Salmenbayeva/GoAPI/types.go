package main

type payment struct {
	PaymentID   string
	Amount     int       `bson:"Amount" json:"Amount"`

}
